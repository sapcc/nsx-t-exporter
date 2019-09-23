package exporter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var managementClusterStates = map[string]float64{
	"STABLE":       1,
	"INITIALIZING": 0,
	"UNSTABLE":     -1,
	"DEGRADED":     -2,
	"UNKNOWN":      -3,
}

var controlClusterStates = map[string]float64{
	"STABLE":         1,
	"NO_CONTROLLERS": 0,
	"UNSTABLE":       -1,
	"DEGRADED":       -2,
	"UNKNOWN":        -3,
}

var nodeConnectivityStates = map[string]float64{
	"CONNECTED": 1,
	"DEGRADED":  0,
	"UNKNOWN":   -1,
}

var transportNodeStates = map[string]float64{
	"SUCCESS":         1,
	"IN_PROGRESS":     0,
	"PENDING":         -1,
	"FAILED":          -2,
	"PARTIAL_SUCCESS": -3,
	"ORPHANED":        -4,
	"UNKNOWN":         -5,
}

var logicalSwitchStates = map[string]float64{
	"SUCCESS":         1,
	"IN_PROGRESS":     0,
	"FAILED":          -1,
	"PARTIAL_SUCCESS": -2,
	"ORPHANED":        -3,
	"UNKNOWN":         -4,
}

var logicalSwitchAdminStates = map[string]float64{
	"UP":   1,
	"DOWN": 0,
}

var noCursor = ""

var (
	endpoints map[string]func(resp *Nsxv3Response, data *Nsxv3Data) string
)

func clusterStatusHandler(resp *Nsxv3Response, data *Nsxv3Data) string {
	var info map[string]interface{}
	json.Unmarshal([]byte(resp.body), &info)

	info = info["mgmt_cluster_status"].(map[string]interface{})

	data.ClusterManagementStatus = managementClusterStates[info["status"].(string)]
	data.ClusterControlStatus = controlClusterStates[info["status"].(string)]

	if onlineNodes, ok := info["online_nodes"]; ok {
		data.ClusterOnlineNodes = float64(len(onlineNodes.([]interface{})))
	}
	if offlineNodes, ok := info["offline_nodes"]; ok {
		data.ClusterOfflineNodes = float64(len(offlineNodes.([]interface{})))
	}

	return noCursor
}

func clusterNodesStatusHandler(resp *Nsxv3Response, data *Nsxv3Data) string {
	var status map[string]interface{}
	json.Unmarshal([]byte(resp.body), &status)

	mgmtNodes := status["management_cluster"].([]interface{})

	for _, node := range mgmtNodes {
		nodeData := new(Nsxv3ManagementNodeData)

		nodeProperties := node.(map[string]interface{})

		nodeData.IP = nodeProperties["role_config"].(map[string]interface{})["api_listen_addr"].(map[string]interface{})["ip_address"].(string)

		nodeData.Connectivity = nodeConnectivityStates[nodeProperties["node_status"].(map[string]interface{})["mgmt_cluster_status"].(map[string]interface{})["mgmt_cluster_status"].(string)]

		nodeData.Version = nodeProperties["node_status"].(map[string]interface{})["version"].(string)

		timeSeries := nodeProperties["node_status_properties"].([]interface{})

		for _, timeSeria := range timeSeries {
			prop := timeSeria.(map[string]interface{})

			load := prop["load_average"].([]interface{})

			nodeData.CPUCores = prop["cpu_cores"].(float64)

			nodeData.LoadAverage[0] = load[0].(float64)
			nodeData.LoadAverage[1] = load[1].(float64)
			nodeData.LoadAverage[2] = load[2].(float64)

			nodeData.MemoryCached = prop["mem_cache"].(float64)
			nodeData.MemoryUse = prop["mem_used"].(float64)
			nodeData.MemoryTotal = prop["mem_total"].(float64)
			nodeData.SwapTotal = prop["swap_total"].(float64)
			nodeData.SwapUse = prop["swap_used"].(float64)

			filesystems := prop["file_systems"].([]interface{})

			for _, filesystem := range filesystems {
				nodeDataStorage := new(Nsxv3NodeStorageData)

				nodeDataStorage.filesystem = filesystem.(map[string]interface{})["mount"].(string)
				nodeDataStorage.totalMetric = float64(filesystem.(map[string]interface{})["total"].(float64))
				nodeDataStorage.usedMetric = float64(filesystem.(map[string]interface{})["used"].(float64))

				nodeData.Storage = append(
					nodeData.Storage,
					*nodeDataStorage)
			}
			// We would like to process only the first of all time serias
			break
		}
		data.ManagementNodes = append(data.ManagementNodes, *nodeData)
	}

	controlNodes := status["controller_cluster"].([]interface{})

	for _, node := range controlNodes {
		nodeData := new(Nsxv3ControlNodeData)

		nodeProperties := node.(map[string]interface{})

		nodeData.IP = nodeProperties["role_config"].(map[string]interface{})["control_plane_listen_addr"].(map[string]interface{})["ip_address"].(string)

		controlNodeStatus := nodeProperties["node_status"].(map[string]interface{})["control_cluster_status"].(map[string]interface{})

		nodeData.Connectivity = nodeConnectivityStates[controlNodeStatus["control_cluster_status"].(string)]
		nodeData.ManagmentConnectivity = nodeConnectivityStates[controlNodeStatus["mgmt_connection_status"].(map[string]interface{})["connectivity_status"].(string)]

		data.ControlNodes = append(data.ControlNodes, *nodeData)
	}
	return noCursor
}

func transportNodeStateHandler(resp *Nsxv3Response, data *Nsxv3Data) string {
	var status map[string]interface{}
	json.Unmarshal([]byte(resp.body), &status)

	nodes := status["results"].([]interface{})

	next := ""
	cursor := status["cursor"]
	if cursor != nil {
		next = cursor.(string)
	}

	for _, node := range nodes {
		nodeData := new(Nsxv3TransportNodeData)

		nodeProperties := node.(map[string]interface{})

		nodeData.ID = nodeProperties["transport_node_id"].(string)
		nodeData.State = transportNodeStates[strings.ToUpper(nodeProperties["state"].(string))]
		nodeData.DeploymentState = transportNodeStates[strings.ToUpper(nodeProperties["node_deployment_state"].(map[string]interface{})["state"].(string))]

		data.TransportNodes = append(data.TransportNodes, *nodeData)
	}
	return next
}

func logicalSwitchAdminStateHander(resp *Nsxv3Response, data *Nsxv3Data) string {
	var status map[string]interface{}
	json.Unmarshal([]byte(resp.body), &status)

	lswitches := status["results"].([]interface{})

	next := ""
	cursor := status["cursor"]
	if cursor != nil {
		next = cursor.(string)
	}

	for _, lswitch := range lswitches {
		lswitchData := new(Nsxv3LogicalSwitchAdminStateData)

		lswitchProperties := lswitch.(map[string]interface{})

		lswitchData.id = lswitchProperties["id"].(string)
		lswitchData.name = lswitchProperties["display_name"].(string)
		lswitchData.adminStateMetric = logicalSwitchAdminStates[lswitchProperties["admin_state"].(string)]

		data.LogicalSwitchesAdminStates = append(data.LogicalSwitchesAdminStates, *lswitchData)
	}
	return next
}

func logicalSwitchStateHander(resp *Nsxv3Response, data *Nsxv3Data) string {
	var status map[string]interface{}
	json.Unmarshal([]byte(resp.body), &status)

	lswitches := status["results"].([]interface{})

	next := ""
	cursor := status["cursor"]
	if cursor != nil {
		next = cursor.(string)
	}

	for _, lswitch := range lswitches {
		lswitchData := new(Nsxv3LogicalSwitchStateData)

		lswitchProperties := lswitch.(map[string]interface{})

		lswitchData.id = lswitchProperties["logical_switch_id"].(string)
		lswitchData.stateMetric = logicalSwitchStates[strings.ToUpper(lswitchProperties["state"].(string))]

		data.LogicalSwitchesStates = append(data.LogicalSwitchesStates, *lswitchData)
	}
	return next
}

func init() {
	endpoints = map[string]func(resp *Nsxv3Response, data *Nsxv3Data) string{
		"/api/v1/cluster/status":         clusterStatusHandler,
		"/api/v1/cluster/nodes/status":   clusterNodesStatusHandler,
		"/api/v1/logical-switches":       logicalSwitchAdminStateHander,
		"/api/v1/logical-switches/state": logicalSwitchStateHander,
		"/api/v1/transport-nodes/state":  transportNodeStateHandler,
	}
}

// gatherData - Collects the data from the API and stores into struct
func (e *Exporter) gatherData(data *Nsxv3Data) error {

	data.ClusterHost = e.NSXv3Configuration.LoginHost

	ch := make(chan *Nsxv3Response, len(endpoints))

	client := GetClient(e.NSXv3Configuration)

	log.Info("Data collection started")

	// Make initial calls
	for path := range endpoints {
		log.Info("Data collection from " + path)
		client.AsyncGetRequest(path, ch)
	}

	cursors := e.gatherDataWithCursors(data, ch)

	// In case data is multypage extract the date with cursors
	for {
		ch = make(chan *Nsxv3Response, len(cursors))

		for path, cursor := range cursors {
			if endpoints[path] != nil {
				log.Info("Data collection from " + path + " with cursor " + cursor)
				client.AsyncGetRequest(path+"?cursor="+cursor, ch)
			}
		}

		// Continue until there are not more cursors
		cursors = e.gatherDataWithCursors(data, ch)

		if len(cursors) <= 0 {
			log.Info("Data collection completed")
			return nil
		}

	}
}

// gatherDataWithCursors - Collects the data from the API and stores into struct
// In case the returned data contains cursor for next page return the map of path/cursors
func (e *Exporter) gatherDataWithCursors(data *Nsxv3Data, ch chan *Nsxv3Response) map[string]string {

	cursors := make(map[string]string)

	if cap(ch) <= 0 {
		return cursors
	}

	reqCountHandled := 0
	for {
		select {
		case resp := <-ch:
			if resp.err != nil {
				log.Errorf("Error scraping NSX-T, Error: %v", resp.err)
				break
			}

			if handler, ok := endpoints[resp.path]; ok {
				cursor := handler(resp, data)
				if cursor != "" {
					cursors[resp.path] = cursor
					log.Info("Data collection completed for " + resp.path + " with cursor " + cursor)
				} else {
					log.Info("Data collection completed for " + resp.path)
				}

			} else {
				log.Error(fmt.Printf("There is no handler function for Path='%s'", resp.path))
			}

			reqCountHandled++

			if reqCountHandled >= cap(ch) {
				log.Info("Data collection completed for bucket with size " + strconv.Itoa(cap(ch)))
				return cursors
			}
		}
	}

}

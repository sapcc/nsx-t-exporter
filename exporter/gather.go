package exporter

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

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
	endpoints map[string]func(resp *Nsxv3Response, data *Nsxv3Data) (string, error)
)

func clusterStatusHandler(resp *Nsxv3Response, data *Nsxv3Data) (string, error) {
	responseStatusCode := resp.response.StatusCode
	if !(responseStatusCode >= 200 && responseStatusCode <= 299) {
		err := errors.New(fmt.Sprintf("Request to endpoint %v returned %d", resp.path, responseStatusCode))
		return noCursor, err
	}

	var info map[string]interface{}
	json.Unmarshal(resp.body, &info)

	info = info["mgmt_cluster_status"].(map[string]interface{})

	data.ClusterManagementStatus = managementClusterStates[info["status"].(string)]
	data.ClusterControlStatus = controlClusterStates[info["status"].(string)]

	if onlineNodes, ok := info["online_nodes"]; ok {
		data.ClusterOnlineNodes = float64(len(onlineNodes.([]interface{})))
	}
	if offlineNodes, ok := info["offline_nodes"]; ok {
		data.ClusterOfflineNodes = float64(len(offlineNodes.([]interface{})))
	}
	return noCursor, nil
}

func clusterNodesStatusHandler(resp *Nsxv3Response, data *Nsxv3Data) (string, error) {

	responseStatusCode := resp.response.StatusCode
	if !(responseStatusCode >= 200 && responseStatusCode <= 299) {
		err := errors.New(fmt.Sprintf("Request to endpoint %v returned %d", resp.path, responseStatusCode))
		return noCursor, err
	}

	var status map[string]interface{}
	json.Unmarshal(resp.body, &status)

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
	return noCursor, nil
}

func transportNodeStateHandler(resp *Nsxv3Response, data *Nsxv3Data) (string, error) {
	responseStatusCode := resp.response.StatusCode
	if !(responseStatusCode >= 200 && responseStatusCode <= 299) {
		err := errors.New(fmt.Sprintf("Request to endpoint %v returned %d", resp.path, responseStatusCode))
		return noCursor, err
	}

	var status map[string]interface{}
	json.Unmarshal(resp.body, &status)

	results := status["results"]
	var nodes []interface{}

	if results != nil {
		nodes = results.([]interface{})
	}

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
	return next, nil
}

func logicalSwitchAdminStateHander(resp *Nsxv3Response, data *Nsxv3Data) (string, error) {
	responseStatusCode := resp.response.StatusCode
	if !(responseStatusCode >= 200 && responseStatusCode <= 299) {
		err := errors.New(fmt.Sprintf("Request to endpoint %v returned %d", resp.path, responseStatusCode))
		return noCursor, err
	}

	var status map[string]interface{}
	json.Unmarshal(resp.body, &status)

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
	return next, nil
}

func logicalSwitchStateHander(resp *Nsxv3Response, data *Nsxv3Data) (string, error) {
	responseStatusCode := resp.response.StatusCode
	if !(responseStatusCode >= 200 && responseStatusCode <= 299) {
		err := errors.New(fmt.Sprintf("Request to endpoint %v returned %d", resp.path, responseStatusCode))
		return noCursor, err
	}

	var status map[string]interface{}
	json.Unmarshal(resp.body, &status)

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
	return next, nil
}

func init() {
	endpoints = map[string]func(resp *Nsxv3Response, data *Nsxv3Data) (string, error){
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
		err := client.AsyncGetRequest(path, ch)
		if err != nil {
			return err
		}
	}
	var gatherDataErrors []interface{}

	cursors, _ := e.gatherDataWithCursors(data, ch, &gatherDataErrors)

	// In case data is multypage extract the date with cursors
	for {
		ch = make(chan *Nsxv3Response, len(cursors))

		for path, cursor := range cursors {
			if endpoints[path] != nil {
				log.Info("Data collection from " + path + " with cursor " + cursor)
				err := client.AsyncGetRequest(path+"?cursor="+cursor, ch)
				if err != nil {
					return err
				}
			}
		}

		// Continue until there are not more cursors
		cursors, _ = e.gatherDataWithCursors(data, ch, &gatherDataErrors)

		if len(cursors) <= 0 && len(gatherDataErrors) == 0 {
			log.Info("Data collection completed")
			data.LastSuccessfulDataFetch = float64(time.Now().Unix())
			return nil
		}

		return errors.New("there were errors while gathering data from endpoints")

	}
}

// gatherDataWithCursors - Collects the data from the API and stores into struct
// In case the returned data contains cursor for next page return the map of path/cursors
func (e *Exporter) gatherDataWithCursors(data *Nsxv3Data, ch chan *Nsxv3Response, gatherDataErrors *[]interface{}) (map[string]string, error) {

	cursors := make(map[string]string)

	if cap(ch) <= 0 {
		return cursors, nil
	}

	reqCountHandled := 0
	for {
		select {
		case resp := <-ch:
			if resp.err != nil {
				log.Errorf("Error scraping NSX-T, Error: %v", resp.err)
				return cursors, resp.err
			}

			if handler, ok := endpoints[resp.path]; ok {
				cursor, err := handler(resp, data)
				if err != nil {
					log.Errorf("Error calling endpoint: %v", err)
					*gatherDataErrors = append(*gatherDataErrors, err)
					reqCountHandled++
					break
				}
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

		}

		if reqCountHandled >= cap(ch) {
			log.Info("Data collection completed for bucket with size " + strconv.Itoa(cap(ch)))
			data.ExtractedActualValues = true
			return cursors, nil
		}
	}

}

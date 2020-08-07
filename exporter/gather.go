package exporter

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

var managementClusterDatabaseStatus = map[string]float64{
	"RUNNING": 0,
	"STOPPED": 1,
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

var logicalPortOperationalStates = map[string]float64{
	"UP":      1,
	"DOWN":    0,
	"UNKNOWN": -1,
}

var noCursor = ""

var (
	endpoints []Nsxv3Resource
)

func getNodeIndexByIP(data *Nsxv3Data, ip string) int {
	for index, element := range data.ManagementNodes {
		if element.IP == ip {
			return index
		}
	}
	return -1
}

func clusterStatusHandler(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	managementClusterInfo := status.state["mgmt_cluster_status"].(map[string]interface{})

	data.ClusterManagementStatus = managementClusterStates[managementClusterInfo["status"].(string)]
	data.ClusterControlStatus = controlClusterStates[managementClusterInfo["status"].(string)]

	if onlineNodes, ok := managementClusterInfo["online_nodes"]; ok {
		data.ClusterOnlineNodes = float64(len(onlineNodes.([]interface{})))
	}
	if offlineNodes, ok := managementClusterInfo["offline_nodes"]; ok {
		data.ClusterOfflineNodes = float64(len(offlineNodes.([]interface{})))
	}
	return noCursor, nil
}

func clusterNodesStatusHandler(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	mgmtNodes := status.state["management_cluster"].([]interface{})
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

	controlNodes := status.state["controller_cluster"].([]interface{})

	for _, node := range controlNodes {
		nodeData := new(Nsxv3ControlNodeData)

		nodeProperties := node.(map[string]interface{})

		nodeData.IP = nodeProperties["role_config"].(map[string]interface{})["control_plane_listen_addr"].(map[string]interface{})["ip_address"].(string)

		controlNodeStatus := nodeProperties["node_status"].(map[string]interface{})["control_cluster_status"].(map[string]interface{})

		nodeData.Connectivity = nodeConnectivityStates[controlNodeStatus["control_cluster_status"].(string)]
		nodeData.ManagementConnectivity = nodeConnectivityStates[controlNodeStatus["mgmt_connection_status"].(map[string]interface{})["connectivity_status"].(string)]

		data.ControlNodes = append(data.ControlNodes, *nodeData)
	}
	return noCursor, nil
}

func managementClusterDatabaseHandler(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	state := status.state["runtime_state"].(string)

	data.DatabaseStatus = managementClusterDatabaseStatus[strings.ToUpper(state)]

	return noCursor, nil
}

func managerNodeFirewallHandler(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	results := status.state["sections_summary"]
	var sectionTypes []interface{}

	if results != nil {
		sectionTypes = results.([]interface{})
	}

	for _, section := range sectionTypes {
		sectionType := section.(map[string]interface{})
		if sectionType["section_type"].(string) == "L3DFW" {
			index := getNodeIndexByIP(data, status.request.URL.Host)
			data.ManagementNodes[index].L3DFWSectionCount = sectionType["section_count"].(float64)
			data.ManagementNodes[index].L3DFWRuleCount = sectionType["rule_count"].(float64)
		}
	}

	return noCursor, nil
}

func managerNodeFirewallSectionsHandler(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	index := getNodeIndexByIP(data, status.request.URL.Host)
	data.ManagementNodes[index].TotalSectionCount = status.state["result_count"].(float64)

	return noCursor, nil
}

func transportNodeStateHandler(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	results := status.state["results"]
	var nodes []interface{}

	if results != nil {
		nodes = results.([]interface{})
	}

	next := noCursor
	cursor := status.state["cursor"]
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

func transportNodesStateHandler(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	data.TransportNodesState.UpCount = status.state["up_count"].(float64)
	data.TransportNodesState.DegradedCount = status.state["degraded_count"].(float64)
	data.TransportNodesState.DownCount = status.state["down_count"].(float64)
	data.TransportNodesState.UnknownCount = status.state["unknown_count"].(float64)
	return noCursor, nil
}

func logicalSwitchAdminStateHander(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	lswitches := status.state["results"].([]interface{})

	next := noCursor
	cursor := status.state["cursor"]
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

func logicalPortsHandler(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	logicalPorts := status.state["results"].([]interface{})

	next := noCursor
	cursor := status.state["cursor"]
	if cursor != nil {
		next = cursor.(string)
	}

	for _, logicalPort := range logicalPorts {
		port := logicalPort.(map[string]interface{})
		logicalPortData := new(Nsxv3LogicalPortOperationalStateData)
		logicalPortData.id = port["id"].(string)
		splittedPortName := strings.Split(port["display_name"].(string), "@")
		logicalPortData.hostID = splittedPortName[len(splittedPortName)-1] // Transport node ID is part of the name
		logicalPortData.operationalStateMetric = logicalPortOperationalStates[port["status"].(map[string]interface{})["status"].(string)]
		data.LogicalPortOperationalStates = append(data.LogicalPortOperationalStates, *logicalPortData)
	}
	return next, nil
}

func logicalSwitchStateHander(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	lswitches := status.state["results"].([]interface{})

	next := noCursor
	cursor := status.state["cursor"]
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

func activityFramworkStatisticsHandler(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	data.Scheduler.TotalQueued = status.state["total_queued"].(float64)
	data.Scheduler.TotalScheduled = status.state["total_scheduled"].(float64)
	data.Scheduler.TotalExecuting = status.state["total_executing"].(float64)
	data.Scheduler.TotalSuspended = status.state["total_suspended"].(float64)
	data.Scheduler.TotalComplete = status.state["total_complete"].(float64)
	return noCursor, nil
}

// endpointHost could be Node FQDN, IP or empty string for NSX-T Load Balancer FQDN/IP
func getEndpointStatus(endpointStatusType Nsxv3ResourceKind, endpointHost string) Nsxv3Resource {
	switch id := endpointStatusType; id {
	case ManagementCluster:
		return Nsxv3Resource{
			kind: endpointStatusType,
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Host: endpointHost, Path: "/api/v1/cluster/status"},
			},
		}
	case ManagementClusterNodes:
		return Nsxv3Resource{
			kind: ManagementClusterNodes,
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Host: endpointHost, Path: "/api/v1/cluster/nodes/status"},
			},
		}
	case ManagerNodeFirewall:
		return Nsxv3Resource{
			kind: ManagerNodeFirewall,
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Host: endpointHost, Path: "/api/v1/firewall/sections/summary"},
			},
		}
	case ManagerNodeFirewallSections:
		return Nsxv3Resource{
			kind: ManagerNodeFirewallSections,
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Host: endpointHost, Path: "/api/v1/firewall/sections"},
			},
		}
	case ManagementClusterDatabase:
		return Nsxv3Resource{
			kind: ManagementClusterDatabase,
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Host: endpointHost, Path: "/api/v1/node/services/datastore/status"},
			},
		}
	case LogicalSwitch:
		return Nsxv3Resource{
			kind: LogicalSwitch,
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Host: endpointHost, Path: "/api/v1/logical-switches"},
			},
		}
	case LogicalSwitchAdmin:
		return Nsxv3Resource{
			kind: LogicalSwitchAdmin,
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Host: endpointHost, Path: "/api/v1/logical-switches/state"},
			},
		}
	case TransportNode:
		return Nsxv3Resource{
			kind: TransportNode,
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Host: endpointHost, Path: "/api/v1/transport-nodes/state"},
			},
		}
	case TransportNodes:
		return Nsxv3Resource{
			kind: TransportNodes,
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Host: endpointHost, Path: "/api/v1/transport-nodes/status"},
			},
		}
	case LogicalPort:
		return Nsxv3Resource{
			kind: LogicalPort,
			request: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Host: endpointHost,
					Path: "/policy/api/v1/search",
					RawQuery: url.Values{
						"query": []string{ // odata query
							strings.Join([]string{
								"resource_type:LogicalPort",
								"_exists_:resource_type",
								"!resource_type:GenericPolicyRealizedResourceORDomain",
								"!_exists_:nsx_id",
							}, " AND ")},
						"page_size":   []string{"100"},
						"data_source": []string{"ALL"},
					}.Encode(),
				},
			},
		}
	case ActivityFrameworkStatistics:
		return Nsxv3Resource{
			kind: ActivityFrameworkStatistics,
			request: &http.Request{
				Method: "GET",
				URL:    &url.URL{Host: endpointHost, Path: "/api/v1/operational/activityframework/scheduler/statistics"},
			},
		}
	}
	return Nsxv3Resource{}
}

func handle(data *Nsxv3Data, status *Nsxv3Resource) (string, error) {
	switch id := status.kind; id {
	case ManagementCluster:
		return clusterStatusHandler(data, status)
	case ManagementClusterNodes:
		return clusterNodesStatusHandler(data, status)
	case ManagementClusterDatabase:
		return managementClusterDatabaseHandler(data, status)
	case ManagerNodeFirewall:
		return managerNodeFirewallHandler(data, status)
	case ManagerNodeFirewallSections:
		return managerNodeFirewallSectionsHandler(data, status)
	case LogicalSwitch:
		return logicalSwitchAdminStateHander(data, status)
	case LogicalSwitchAdmin:
		return logicalSwitchStateHander(data, status)
	case TransportNode:
		return transportNodeStateHandler(data, status)
	case TransportNodes:
		return transportNodesStateHandler(data, status)
	case LogicalPort:
		return logicalPortsHandler(data, status)
	case ActivityFrameworkStatistics:
		return activityFramworkStatisticsHandler(data, status)
	}
	return noCursor, fmt.Errorf("Unsupported Endpoint Type %v", status.kind)
}

func (e *Exporter) gatherWave(data *Nsxv3Data, endpoints []Nsxv3Resource) error {
	clients := map[string]Nsxv3Client{
		e.NSXv3Configuration.LoginHost: GetClient(e.NSXv3Configuration),
	}

	chSize := len(endpoints)

	ch := make(chan error, chSize)

	for id := range endpoints {
		host := endpoints[id].request.URL.Host
		if host == "" {
			host = e.NSXv3Configuration.LoginHost
		}
		client, ok := clients[host]
		if !ok {
			config := e.NSXv3Configuration
			config.LoginHost = host
			client = GetClient(config)
			clients[host] = client
		}
		go e.updateData(data, &client, &endpoints[id], ch)
	}

	var errs []interface{}

	for i := 0; i < chSize; i++ {
		select {
		case err := <-ch:
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) != 0 {
		e := errors.New("Data collection completed with errors")
		for _, err := range errs {
			log.Error(err)
		}
		return e
	}

	return nil
}

func (e *Exporter) gather(data *Nsxv3Data) error {
	log.Info("Data collection started")
	data.ClusterHost = e.NSXv3Configuration.LoginHost

	var err error

	err = e.gatherWave(
		data,
		[]Nsxv3Resource{
			getEndpointStatus(ManagementCluster, ""),
			getEndpointStatus(ManagementClusterNodes, ""),
			getEndpointStatus(ManagementClusterDatabase, ""),
			getEndpointStatus(LogicalSwitchAdmin, ""),
			getEndpointStatus(LogicalSwitch, ""),
			getEndpointStatus(TransportNode, ""),
			getEndpointStatus(TransportNodes, ""),
			getEndpointStatus(LogicalPort, ""),
			getEndpointStatus(ActivityFrameworkStatistics, ""),
		})

	if err != nil {
		return err
	}

	endpoints := []Nsxv3Resource{}
	for _, element := range data.ManagementNodes {
		endpoints = append(endpoints,
			getEndpointStatus(ManagerNodeFirewall, element.IP),
			getEndpointStatus(ManagerNodeFirewallSections, element.IP))
	}

	err = e.gatherWave(data, endpoints)

	if err != nil {
		return err
	}

	data.ExtractedActualValues = true
	data.LastSuccessfulDataFetch = float64(time.Now().Unix())

	log.Info("Data collection completed")
	return nil
}

func (e *Exporter) updateData(data *Nsxv3Data, client *Nsxv3Client, status *Nsxv3Resource, ch chan error) {
	client.updateEndpointStatus(status)

	if status.err != nil {
		ch <- status.err
		return
	}

	cursor, err := handle(data, status)

	if err == nil {
		if cursor == noCursor {
			ch <- nil
			return
		}

		nextStatus := getEndpointStatus(status.kind, status.request.URL.Host)

		query, _ := url.ParseQuery(nextStatus.request.URL.RawQuery)
		query.Add("cursor", cursor)
		nextStatus.request.URL.RawQuery = query.Encode()

		e.updateData(data, client, &nextStatus, ch)
	} else {
		ch <- err
	}
}

package exporter

import "github.com/prometheus/client_golang/prometheus"

// GetMetricsDescription - creates Prometheus metrics description
func GetMetricsDescription() map[string]*prometheus.Desc {

	APIMetrics := make(map[string]*prometheus.Desc)

	APIMetrics["ManagementClusterStatus"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster_management", "status"),
		"NSX-T management cluster status - STABLE=1, INITIALIZING=0, UNSTABLE=-1, DEGRADED=-2, UNKNOWN=-3",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["ControlClusterStatus"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster_control", "status"),
		"NSX-T control cluster status - STABLE=1, NO_CONTROLLERS=0, UNSTABLE=-1, DEGRADED=-2, UNKNOWN=-3",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["ManagementClusterNodesOnline"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster_management", "online_nodes"),
		"NSX-T management cluster online nodes",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["ManagementClusterNodesOffline"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster_management", "offline_nodes"),
		"NSX-T management cluster offline nodes",
		[]string{"nsxv3_manager_hostname"}, nil,
	)

	APIMetrics["ManagementNodeConnectivity"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "connectivity"),
		"NSX-T management node connectivity - CONNECTED > 0, DISCONNECTED = 0, UNKNOWN < 0",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeCpuCores"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "cpu_cores"),
		"NSX-T management node cpu cores",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeLoadAverage"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "load_average"),
		"NSX-T management node average load",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip", "minutes"}, nil,
	)

	APIMetrics["ManagementNodeMemoryUse"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "memory_use"),
		"NSX-T management node memory use",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeMemoryTotal"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "memory_total"),
		"NSX-T management node memory total",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeMemoryCached"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "memory_cached"),
		"NSX-T management node cached memory",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeSwapUse"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "swap_use"),
		"NSX-T management node swap use",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeSwapTotal"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "swap_total"),
		"NSX-T management node swap total",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ManagementNodeStorageUse"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "storage_use"),
		"NSX-T management node storage use",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip", "filesystem"}, nil,
	)

	APIMetrics["ManagementNodeStorageTotal"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "management_node", "storage_total"),
		"NSX-T management node storage total",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip", "filesystem"}, nil,
	)

	APIMetrics["ControlNodeConnectivity"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "control_node", "connectivity"),
		"NSX-T control node connectivity - CONNECTED=1, DISCONNECTED=0, UNKNOWN=-1",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["ControlNodeManagmentConnectivity"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "control_node", "management_connectivity"),
		"NSX-T control node management connectivity - CONNECTED > 0, DISCONNECTED = 0, UNKNOWN < 0",
		[]string{"nsxv3_manager_hostname", "nsxv3_node_ip"}, nil,
	)

	APIMetrics["LogicalSwitchAdminStateUp"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "logical_switch", "admin_state"),
		"NSX-T logical switch admin state - UP=1, DOWN=0",
		[]string{"nsxv3_manager_hostname", "name"}, nil,
	)

	return APIMetrics
}

// processMetrics - processes the response data and sets the metrics using it as a source
func (e *Exporter) processMetrics(data *Nsxv3Data, ch chan<- prometheus.Metric) error {

	// Prometheus scrape metric callback (concurrent)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementClusterStatus"],
		prometheus.GaugeValue,
		data.ClusterManagementStatus,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ControlClusterStatus"],
		prometheus.GaugeValue,
		data.ClusterControlStatus,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementClusterNodesOnline"],
		prometheus.GaugeValue,
		data.ClusterOnlineNodes,
		data.ClusterHost)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementClusterNodesOffline"],
		prometheus.GaugeValue,
		data.ClusterOfflineNodes,
		data.ClusterHost)

	for _, element := range data.ManagementNodes {
		e.processManagementNodeMetrics(data.ClusterHost, &element, ch)
	}

	for _, element := range data.ControlNodes {
		e.processControlNodeMetrics(data.ClusterHost, &element, ch)
	}

	for _, element := range data.LogicalSwitches {
		e.processLogicalSwitchMetrics(data.ClusterHost, &element, ch)
	}

	return nil

}

func (e *Exporter) processManagementNodeMetrics(host string, data *Nsxv3ManagementNodeData, ch chan<- prometheus.Metric) error {
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeConnectivity"],
		prometheus.GaugeValue,
		data.Connectivity,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeCpuCores"],
		prometheus.GaugeValue,
		data.CPUCores,
		host, data.IP)

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeLoadAverage"],
		prometheus.GaugeValue,
		data.LoadAverage[0],
		host, data.IP, "1") // 1 minute
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeLoadAverage"],
		prometheus.GaugeValue,
		data.LoadAverage[1],
		host, data.IP, "5") // 5 minutes
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeLoadAverage"],
		prometheus.GaugeValue,
		data.LoadAverage[2],
		host, data.IP, "15") // 15 minutes

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeMemoryUse"],
		prometheus.GaugeValue,
		data.MemoryUse,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeMemoryTotal"],
		prometheus.GaugeValue,
		data.MemoryTotal,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeMemoryCached"],
		prometheus.GaugeValue,
		data.MemoryCached,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeSwapUse"],
		prometheus.GaugeValue,
		data.SwapUse,
		host, data.IP)
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ManagementNodeSwapTotal"],
		prometheus.GaugeValue,
		data.SwapTotal,
		host, data.IP)

	for _, element := range data.Storage {
		ch <- prometheus.MustNewConstMetric(
			e.APIMetrics["ManagementNodeStorageTotal"],
			prometheus.GaugeValue,
			element.totalMetric,
			host, data.IP, element.filesystem)
		ch <- prometheus.MustNewConstMetric(
			e.APIMetrics["ManagementNodeStorageUse"],
			prometheus.GaugeValue,
			element.usedMetric,
			host, data.IP, element.filesystem)
	}

	return nil
}

func (e *Exporter) processControlNodeMetrics(host string, data *Nsxv3ControlNodeData, ch chan<- prometheus.Metric) error {

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ControlNodeConnectivity"],
		prometheus.GaugeValue,
		data.Connectivity,
		host, data.IP)

	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["ControlNodeManagmentConnectivity"],
		prometheus.GaugeValue,
		data.ManagmentConnectivity,
		host, data.IP)

	return nil
}

func (e *Exporter) processLogicalSwitchMetrics(host string, data *Nsxv3LogicalSwitchData, ch chan<- prometheus.Metric) error {
	ch <- prometheus.MustNewConstMetric(
		e.APIMetrics["LogicalSwitchAdminStateUp"],
		prometheus.GaugeValue,
		data.statusMetric,
		host, data.name)
	return nil
}

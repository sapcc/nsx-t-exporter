package exporter

import "github.com/prometheus/client_golang/prometheus"

// GetMetricsDescription - creates Prometheus metrics description
func GetMetricsDescription() map[string]*prometheus.Desc {

	APIMetrics := make(map[string]*prometheus.Desc)

	APIMetrics["ManagementClusterActive"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster", "management_active"),
		"NSX-T management cluster status active != 0",
		[]string{"nsxv3_manager"}, nil,
	)

	APIMetrics["ControlClusterActive"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster", "control_active"),
		"NSX-T control cluster status active != 0",
		[]string{"nsxv3_manager"}, nil,
	)

	APIMetrics["ClusterNodesOnline"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster", "online_nodes"),
		"NSX-T cluster online nodes",
		[]string{"nsxv3_manager"}, nil,
	)

	APIMetrics["ClusterNodesOffline"] = prometheus.NewDesc(
		prometheus.BuildFQName("nsxv3", "cluster", "offline_nodes"),
		"NSX-T cluster offline nodes",
		[]string{"nsxv3_manager"}, nil,
	)

	return APIMetrics
}

// processMetrics - processes the response data and sets the metrics using it as a source
func (e *Exporter) processMetrics(data *Nsxv3Data, ch chan<- prometheus.Metric) error {

	// Prometheus scrape metric callback (concurrent)
	ch <- prometheus.MustNewConstMetric(e.APIMetrics["ManagementClusterActive"], prometheus.GaugeValue, data.Nsxv3ClusterManagementActive, data.Nsxv3ClusterHost)
	ch <- prometheus.MustNewConstMetric(e.APIMetrics["ControlClusterActive"], prometheus.GaugeValue, data.Nsxv3ClusterControlActive, data.Nsxv3ClusterHost)
	ch <- prometheus.MustNewConstMetric(e.APIMetrics["ClusterNodesOnline"], prometheus.GaugeValue, data.Nsxv3ClusterOnlineNodes, data.Nsxv3ClusterHost)
	ch <- prometheus.MustNewConstMetric(e.APIMetrics["ClusterNodesOffline"], prometheus.GaugeValue, data.Nsxv3ClusterOfflineNodes, data.Nsxv3ClusterHost)

	return nil

}

package exporter

import (
	"github.com/sapcc/nsx-t-exporter/config"

	"github.com/prometheus/client_golang/prometheus"
)

// Exporter is used to store Metrics data and embeds the config struct.
// This is done so that the relevant functions have easy access to the
// user defined runtime configuration when the Collect method is called.
type Exporter struct {
	APIMetrics map[string]*prometheus.Desc
	config.NSXv3Configuration
}

// Nsxv3Data represent current snapshot of metrics
type Nsxv3Data struct {
	Nsxv3ClusterHost string

	Nsxv3ClusterManagementActive float64
	Nsxv3ClusterControlActive    float64
	Nsxv3ClusterOnlineNodes      float64
	Nsxv3ClusterOfflineNodes     float64

	NsxV3FirewallRulesTotal    float64
	NsxV3IPSetsTotal           float64
	NsxV3INSGroupsTotal        float64
	NsxV3ILogicalSwitchesTotal float64
}

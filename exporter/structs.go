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

// Nsxv3LogicalSwitchData represent the current snapshot of metrics
// for a logical switch
type Nsxv3LogicalSwitchData struct {
	statusMetric float64
	name         string
}

// Nsxv3NodeStorageData represent the current snapshot of metrics
// for a single node filesystem
type Nsxv3NodeStorageData struct {
	usedMetric  float64
	totalMetric float64
	filesystem  string
}

// Nsxv3ManagementNodeData represent the current snapshot of metrics for a single management node
type Nsxv3ManagementNodeData struct {
	IP           string
	Connectivity float64
	CPUCores     float64
	// Averadge load index [0] := 1min, [1] := 5min, [2] := 15min
	LoadAverage  [3]float64
	MemoryUse    float64
	MemoryTotal  float64
	MemoryCached float64
	SwapUse      float64
	SwapTotal    float64
	Storage      []Nsxv3NodeStorageData
}

// Nsxv3ControlNodeData represent the current snapshot of metrics for a single control node
type Nsxv3ControlNodeData struct {
	IP                    string
	Connectivity          float64
	ManagmentConnectivity float64
}

// Nsxv3Data represent the current snapshot of metrics
type Nsxv3Data struct {
	ClusterHost             string
	ClusterManagementStatus float64
	ClusterControlStatus    float64
	ClusterOnlineNodes      float64
	ClusterOfflineNodes     float64
	ManagementNodes         []Nsxv3ManagementNodeData
	ControlNodes            []Nsxv3ControlNodeData
	LogicalSwitches         []Nsxv3LogicalSwitchData
}

package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sapcc/nsx-t-exporter/config"
)

// Exporter is used to store Metrics data and embeds the config struct.
// This is done so that the relevant functions have easy access to the
// user defined runtime configuration when the Collect method is called.
type Exporter struct {
	Cache      Nsxv3Data
	APIMetrics map[string]*prometheus.Desc
	config.NSXv3Configuration
}

// Nsxv3LogicalSwitchAdminStateData represent the current snapshot of metrics
// for a logical switch admin state
type Nsxv3LogicalSwitchAdminStateData struct {
	adminStateMetric float64
	name             string
	id               string
}

// Nsxv3LogicalPortOperationalStateData represent the current snapshot of metrics
// for a logical port operational state
type Nsxv3LogicalPortOperationalStateData struct {
	operationalStateMetric float64
	id                     string
	hostID                 string // Transport Node ID
}

// Nsxv3LogicalSwitchStateData represent the current snapshot of metrics
// for a logical switch state
type Nsxv3LogicalSwitchStateData struct {
	stateMetric float64
	id          string
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
	// Average load index [0] := 1min, [1] := 5min, [2] := 15min
	LoadAverage  [3]float64
	MemoryUse    float64
	MemoryTotal  float64
	MemoryCached float64
	SwapUse      float64
	SwapTotal    float64
	Storage      []Nsxv3NodeStorageData
	Version      string
}

// Nsxv3ControlNodeData represent the current snapshot of metrics for a single control node
type Nsxv3ControlNodeData struct {
	IP                     string
	Connectivity           float64
	ManagementConnectivity float64
}

// Nsxv3TransportNodeData represent the current snapshot of metrics for a single transport node
type Nsxv3TransportNodeData struct {
	ID              string
	State           float64
	DeploymentState float64
}

// Nsxv3TransportNodesStateData represent the current snapshot of the transport nodes state
type Nsxv3TransportNodesStateData struct {
	UpCount       float64
	DegradedCount float64
	DownCount     float64
	UnknownCount  float64
}

// Nsxv3Data represent the current snapshot of metrics
type Nsxv3Data struct {
	ClusterHost                  string
	ClusterManagementStatus      float64
	ClusterControlStatus         float64
	ClusterOnlineNodes           float64
	ClusterOfflineNodes          float64
	DatabaseStatus               float64
	ManagementNodes              []Nsxv3ManagementNodeData
	ControlNodes                 []Nsxv3ControlNodeData
	TransportNodes               []Nsxv3TransportNodeData
	TransportNodesState          Nsxv3TransportNodesStateData
	LogicalSwitchesAdminStates   []Nsxv3LogicalSwitchAdminStateData
	LogicalSwitchesStates        []Nsxv3LogicalSwitchStateData
	ExtractedActualValues        bool
	LastSuccessfulDataFetch      float64
	LogicalPortOperationalStates []Nsxv3LogicalPortOperationalStateData
}

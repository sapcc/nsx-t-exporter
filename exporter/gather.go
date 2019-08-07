package exporter

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type managementClusterStatus struct {
	Status       string                        `json:"status"`
	OnlineNodes  []map[string]*json.RawMessage `json:"online_nodes"`
	OfflineNodes []map[string]*json.RawMessage `json:"offline_nodes"`
}

type controlClusterStatus struct {
	Status string `json:"status"`
}

type clusterStatus struct {
	ManagementClusterStatus managementClusterStatus `json:"mgmt_cluster_status"`
	ControlClusterStatus    controlClusterStatus    `json:"control_cluster_status"`
}

const httpPathClusterStatus = "/api/v1/cluster/status"

var (
	endpoints map[string]func(resp *Nsxv3Response, data *Nsxv3Data)
)

func clusterStatusHandler(resp *Nsxv3Response, data *Nsxv3Data) {
	var status clusterStatus
	json.Unmarshal([]byte(resp.body), &status)
	if status.ManagementClusterStatus.Status == "STABLE" {
		data.Nsxv3ClusterManagementActive = 1
	}
	if status.ControlClusterStatus.Status == "STABLE" {
		data.Nsxv3ClusterControlActive = 1
	}
	data.Nsxv3ClusterOnlineNodes = float64(len(status.ManagementClusterStatus.OnlineNodes))
	data.Nsxv3ClusterOfflineNodes = float64(len(status.ManagementClusterStatus.OfflineNodes))
}

func init() {
	endpoints = map[string]func(resp *Nsxv3Response, data *Nsxv3Data){
		"/api/v1/cluster/status": clusterStatusHandler,
	}
}

// gatherData - Collects the data from the API and stores into struct
func (e *Exporter) gatherData(data *Nsxv3Data) error {

	data.Nsxv3ClusterHost = e.NSXv3Configuration.LoginHost

	ch := make(chan *Nsxv3Response, len(endpoints))

	client := GetClient(e.NSXv3Configuration)

	client.AsyncGetRequest(httpPathClusterStatus, ch)

	reqCountHandled := 0
	for {
		select {
		case resp := <-ch:
			if resp.err != nil {
				log.Errorf("Error scraping NSX-T, Error: %v", resp.err)
				break
			}

			if handler, ok := endpoints[resp.path]; ok {
				handler(resp, data)
			}else{
				log.Error(fmt.Printf("There is no handler function for Path='%s'", resp.path))
			}

			reqCountHandled++

			if reqCountHandled == len(endpoints) {
				return nil
			}
		}

	}

}

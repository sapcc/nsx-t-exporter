package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// Describe the API metrics in Prometheus
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	for _, m := range e.APIMetrics {
		ch <- m
	}

}

// Collect is a Prometheus callback for scrape operations (/metrics page)
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	data := Nsxv3Data{}

	err := e.gatherData(&data)
	if err != nil {
		return
	}

	if err != nil {
		log.Errorf("Error gathering Data from remote API: %v", err)
		return
	}

	// Set Prometheus gauge metrics using the data
	err = e.processMetrics(&data, ch)

	if err != nil {
		log.Error("Error Processing Metrics", err)
		return
	}

	log.Info("All Metrics successfully collected")

}

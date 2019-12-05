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

// CollectAsync - Asynchronously scraps the API and stores into struct
func (e *Exporter) CollectAsync() {

	data := Nsxv3Data{}

	err := e.gatherData(&data)

	// in case of error, keep the previously acquired data (cache)
	// otherwise update the existing cache
	if err != nil {
		log.Errorf("Error gathering Data from remote API: %v", err)
		return
	}
	e.Cache = data
}

// Collect is a Prometheus callback for scrape operations (/metrics page)
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	// Set Prometheus gauge metrics using the e.Cache
	err := e.processMetrics(&e.Cache, ch)

	if err != nil {
		log.Error("Error Processing Metrics", err)
		return
	}

	log.Info("All Metrics successfully collected")

}

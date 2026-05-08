// SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sapcc/nsx-t-exporter/config"
	"github.com/sapcc/nsx-t-exporter/exporter"

	"github.com/fatih/structs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	"github.com/sapcc/go-api-declarations/bininfo"
)

var (
	indexPage = `<html>
	<head><title>NSX-T Exporter</title></head>
	<body>
		<h1>NSX-T Prometheus Metrics Exporter</h1>
		<p>For more information, visit <a href="https://github.com/sapcc/nsxv3-exporter">GitHub</a></p>
		<p><a href="/metrics">Metrics</a></p>
	</body>
</html>`
)

var (
	exporterConfig config.NSXv3Configuration
	metrics        map[string]*prometheus.Desc
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	bininfo.HandleVersionArgument()

	exporterConfig = config.Init()

	metrics = exporter.GetMetricsDescription()
}

func main() {
	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	// A common pattern is to re-use fields between logging statements
	loggingContext := exporterConfig
	loggingContext.LoginPassword = "*******"
	contextLogger := log.WithFields(structs.Map(loggingContext))

	contextLogger.Info("Starting Exporter")

	ex := exporter.Exporter{
		APIMetrics:         metrics,
		NSXv3Configuration: exporterConfig,
	}

	// Async scrap, required to reduce response time of /metrics API
	go func() {
		for {
			start := time.Now()
			ex.CollectAsync()
			elapsed := time.Since(start)
			schedule := time.Duration(exporterConfig.ScrapScheduleSeconds) * time.Second
			time.Sleep(schedule - elapsed)
		}
	}()

	// Register Metrics from each of the endpoints
	// This invokes the Collect method through the prometheus client libraries.
	prometheus.MustRegister(&ex)

	// Setup HTTP handler
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(indexPage)); err != nil {
			log.Error(err)
		}
	})

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", exporterConfig.ScrapPort),
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

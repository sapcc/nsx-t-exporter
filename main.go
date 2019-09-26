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

	exporter := exporter.Exporter{
		APIMetrics:         metrics,
		NSXv3Configuration: exporterConfig,
	}

	// Async scrap, required to reduce response time of /metrics API
	go func() {
		for {
			start := time.Now()
			exporter.CollectAsync()
			elapsed := time.Now().Sub(start)
			schedule := time.Duration(exporterConfig.ScrapScheduleSeconds) * time.Second
			time.Sleep(schedule - elapsed)
		}
	}()

	// Register Metrics from each of the endpoints
	// This invokes the Collect method through the prometheus client libraries.
	prometheus.MustRegister(&exporter)

	// Setup HTTP handler
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(indexPage))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", exporterConfig.ScrapPort), nil))
}

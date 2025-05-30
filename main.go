package main

import (
	"github.com/amirhnajafiz/k8sese/internal/collector"
	"github.com/amirhnajafiz/k8sese/internal/metrics"
)

func main() {
	// create a new metrics instance
	mtx, err := metrics.NewMetrics()
	if err != nil {
		panic(err)
	}

	// start the metrics server on port 8080
	metrics.StartMetricsServer(8080)

	// continue to collect metrics
	collector.Start(mtx)
}

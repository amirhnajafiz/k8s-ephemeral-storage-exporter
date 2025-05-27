package main

import (
	"github.com/amirhnajafiz/k8sese/internal/metrics"
	"github.com/amirhnajafiz/k8sese/internal/runner"
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
	runner.Start(mtx)
}

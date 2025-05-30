package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.uber.org/zap"
)

// StartMetricsServer starts an HTTP server that serves Prometheus metrics.
func StartMetricsServer(logr *zap.Logger, port int) {
	go func() {
		// create a new HTTP server
		http.Handle("/metrics", promhttp.Handler())
		addr := fmt.Sprintf(":%d", port)

		logr.Info("starting metrics server", zap.String("address", addr))

		// start the server
		if err := http.ListenAndServe(addr, nil); err != nil {
			logr.Fatal("failed to start metrics server", zap.Error(err))
		}
	}()
}

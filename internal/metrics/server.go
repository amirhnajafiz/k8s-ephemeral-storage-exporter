package metrics

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartMetricsServer starts an HTTP server that serves Prometheus metrics.
func StartMetricsServer(port int) {
	go func() {
		// create a new HTTP server
		http.Handle("/metrics", promhttp.Handler())
		addr := fmt.Sprintf(":%d", port)
		log.Printf("starting metrics server on %s", addr)

		// start the server
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("failed to start metrics server: %v", err)
		}
	}()
}

package metrics

import "github.com/prometheus/client_golang/prometheus"

// constant values for namespace and subsystem
const (
	NS = "KES"
	SS = "exporter"
)

// Metrics holds the Prometheus metrics for the exporter.
type Metrics struct {
	storageUsageBytes *prometheus.GaugeVec
}

// NewMetrics initializes and registers the Prometheus metrics for the exporter.
func NewMetrics() (*Metrics, error) {
	// create Prometheus metrics
	sub := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "ephemeral_storage_usage_bytes",
			Help:      "Ephemeral storage usage in bytes",
		},
		[]string{"pod", "namespace", "node", "container"},
	)

	// register the metrics with Prometheus
	if err := prometheus.Register(sub); err != nil {
		return nil, err
	}

	return &Metrics{
		storageUsageBytes: sub,
	}, nil
}

// SetStorageUsageBytes sets the ephemeral storage usage in bytes for a specific pod, namespace, node, and container.
func (m *Metrics) SetStorageUsageBytes(pod, namespace, node, container string, value float64) {
	m.storageUsageBytes.WithLabelValues(pod, namespace, node, container).Set(value)
}

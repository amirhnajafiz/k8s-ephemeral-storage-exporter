package metrics

import "github.com/prometheus/client_golang/prometheus"

// constant values for namespace and subsystem
const (
	NS = "KES"
	SS = "exporter"
)

// Metrics holds the Prometheus metrics for the exporter.
type Metrics struct {
	ephemeralStorageUsageBytes *prometheus.GaugeVec
	containerStorageUsageBytes *prometheus.GaugeVec
}

// NewMetrics initializes and registers the Prometheus metrics for the exporter.
func NewMetrics() (*Metrics, error) {
	// create Prometheus metrics
	esub := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "ephemeral_storage_usage_bytes",
			Help:      "Ephemeral storage usage in bytes",
		},
		[]string{"pod", "namespace", "node"},
	)
	cmub := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_storage_usage_bytes",
			Help:      "Container storage (memory, logs, and roofs) usage in bytes",
		},
		[]string{"pod", "namespace", "node", "container", "type"},
	)

	// register the metrics with Prometheus
	if err := prometheus.Register(esub); err != nil {
		return nil, err
	}
	if err := prometheus.Register(cmub); err != nil {
		return nil, err
	}

	return &Metrics{
		ephemeralStorageUsageBytes: esub,
		containerStorageUsageBytes: cmub,
	}, nil
}

// SetStorageUsageBytes sets the ephemeral storage usage in bytes for a specific pod, namespace, and node.
func (m *Metrics) SetStorageUsageBytes(pod, namespace, node string, value float64) {
	m.ephemeralStorageUsageBytes.WithLabelValues(pod, namespace, node).Set(value)
}

// SetContainerStorageUsageBytes sets the storage usage in bytes for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerStorageUsageBytes(pod, namespace, node, container, typ string, value float64) {
	m.containerStorageUsageBytes.WithLabelValues(pod, namespace, node, container, typ).Set(value)
}

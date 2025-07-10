package metrics

// SetStorageUsageBytes sets the ephemeral storage usage in bytes for a specific pod, namespace, and node.
func (m *Metrics) SetStorageUsageBytes(pod, namespace, node string, value float64) {
	m.ephemeralStorageUsageBytes.WithLabelValues(pod, namespace, node).Set(value)
}

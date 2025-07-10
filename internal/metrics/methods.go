package metrics

// SetEphemeralStorageValues sets the ephemeral storage metrics for a specific pod, namespace, and node.
func (m *Metrics) SetEphemeralStorageValues(
	pod, namespace, node, uid string,
	used, available, capacity float64,
) {
	m.ephemeralStorageUsageBytes.WithLabelValues(pod, namespace, node, uid).Set(used)
	m.ephemeralStorageAvailableBytes.WithLabelValues(pod, namespace, node, uid).Set(available)
	m.ephemeralStorageCapacityBytes.WithLabelValues(pod, namespace, node, uid).Set(capacity)
}

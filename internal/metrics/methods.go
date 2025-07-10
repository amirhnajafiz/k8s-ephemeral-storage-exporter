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

// SetEphemeralStorageInodes sets the ephemeral storage inode metrics for a specific pod, namespace, and node.
func (m *Metrics) SetEphemeralStorageInodes(
	pod, namespace, node, uid string,
	used, available, capacity float64,
) {
	m.ephemeralStorageInodesFree.WithLabelValues(pod, namespace, node, uid).Set(available)
	m.ephemeralStorageInodesUsed.WithLabelValues(pod, namespace, node, uid).Set(used)
	m.ephemeralStorageInodes.WithLabelValues(pod, namespace, node, uid).Set(capacity)
}

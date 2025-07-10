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
	m.ephemeralStorageInodesUsed.WithLabelValues(pod, namespace, node, uid).Set(used)
	m.ephemeralStorageInodesFree.WithLabelValues(pod, namespace, node, uid).Set(available)
	m.ephemeralStorageInodes.WithLabelValues(pod, namespace, node, uid).Set(capacity)
}

// SetContainerMemoryValues sets the memory metrics for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerMemoryValues(
	pod, namespace, node, uid, container string,
	used, available, capacity float64,
) {
	m.containerMemoryUsageBytes.WithLabelValues(pod, namespace, node, uid, container).Set(used)
	m.containerMemoryAvailableBytes.WithLabelValues(pod, namespace, node, uid, container).Set(available)
	m.containerMemoryCapacityBytes.WithLabelValues(pod, namespace, node, uid, container).Set(capacity)
}

// SetContainerRootfsValues sets the root filesystem metrics for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerRootfsValues(
	pod, namespace, node, uid, container string,
	used, available, capacity float64,
) {
	m.containerRootfsUsageBytes.WithLabelValues(pod, namespace, node, uid, container).Set(used)
	m.containerRootfsAvailableBytes.WithLabelValues(pod, namespace, node, uid, container).Set(available)
	m.containerRootfsCapacityBytes.WithLabelValues(pod, namespace, node, uid, container).Set(capacity)
}

// SetContainerRootfsInodes sets the root filesystem inode metrics for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerRootfsInodes(
	pod, namespace, node, uid, container string,
	used, available, capacity float64,
) {
	m.containerRootfsInodesUsed.WithLabelValues(pod, namespace, node, uid, container).Set(used)
	m.containerRootfsInodesFree.WithLabelValues(pod, namespace, node, uid, container).Set(available)
	m.containerRootfsInodes.WithLabelValues(pod, namespace, node, uid, container).Set(capacity)
}

// SetContainerLogsValues sets the logs metrics for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerLogsValues(
	pod, namespace, node, uid, container string,
	used, available, capacity float64,
) {
	m.containerLogsUsageBytes.WithLabelValues(pod, namespace, node, uid, container).Set(used)
	m.containerLogsAvailableBytes.WithLabelValues(pod, namespace, node, uid, container).Set(available)
	m.containerLogsCapacityBytes.WithLabelValues(pod, namespace, node, uid, container).Set(capacity)
}

// SetContainerLogsInodes sets the logs inode metrics for a specific container in a pod, namespace, and node.
func (m *Metrics) SetContainerLogsInodes(
	pod, namespace, node, uid, container string,
	used, available, capacity float64,
) {
	m.containerLogsInodesUsed.WithLabelValues(pod, namespace, node, uid, container).Set(used)
	m.containerLogsInodesFree.WithLabelValues(pod, namespace, node, uid, container).Set(available)
	m.containerLogsInodes.WithLabelValues(pod, namespace, node, uid, container).Set(capacity)
}

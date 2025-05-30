package collector

import (
	"github.com/amirhnajafiz/k8sese/internal/metrics"
	"github.com/amirhnajafiz/k8sese/pkg/fetch"
	"github.com/amirhnajafiz/k8sese/pkg/types"
)

// Start initiates the process of fetching storage usage metrics from the kubelet summary endpoint
// and updates the provided metrics instance with the data.
func Start(m *metrics.Metrics) {
	for {
		// build the HTTP request to the kubelet summary endpoint
		req, err := buildHTTPRequest()
		if err != nil {
			// handle error (e.g., log it, retry, etc.)
			continue
		}

		// perform the HTTP GET request
		resp, err := fetch.GET(req)
		if err != nil {
			// handle error (e.g., log it, retry, etc.)
			continue
		}

		// decode the JSON response into a summary structure
		var summary types.Summary
		if err := fetch.JSON(resp, &summary); err != nil {
			// handle error (e.g., log it, retry, etc.)
			continue
		}

		// process the summary data and update the metrics
		for _, pod := range summary.Pods {
			m.SetStorageUsageBytes(
				pod.PodRef.Name,
				pod.PodRef.Namespace,
				summary.Node.NodeName,
				float64(pod.EphemeralStorage.UsedBytes),
			)

			for _, container := range pod.Containers {
				m.SetContainerStorageUsageBytes(
					pod.PodRef.Name,
					pod.PodRef.Namespace,
					summary.Node.NodeName,
					container.Name,
					"memory",
					float64(container.Memory.UsageBytes),
				)
				m.SetContainerStorageUsageBytes(
					pod.PodRef.Name,
					pod.PodRef.Namespace,
					summary.Node.NodeName,
					container.Name,
					"logs",
					float64(container.Logs.UsedBytes),
				)
				m.SetContainerStorageUsageBytes(
					pod.PodRef.Name,
					pod.PodRef.Namespace,
					summary.Node.NodeName,
					container.Name,
					"rootfs",
					float64(container.Rootfs.UsedBytes),
				)
			}
		}
	}
}

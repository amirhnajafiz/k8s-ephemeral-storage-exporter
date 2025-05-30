package collector

import (
	"github.com/amirhnajafiz/k8sese/internal/metrics"
	"github.com/amirhnajafiz/k8sese/pkg/fetch"
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
		var summary summary
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
				"c",
				float64(pod.EphemeralStorage.UsedBytes),
			)
		}
	}
}

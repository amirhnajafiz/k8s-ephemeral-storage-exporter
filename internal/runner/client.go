package runner

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/amirhnajafiz/k8sese/internal/metrics"
)

func collectMetrics(m *metrics.Metrics) {
	resp, err := http.Get(kubeletSummaryEndpoint)
	if err != nil {
		log.Println("Error fetching stats:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var summary summary
	json.Unmarshal(body, &summary)

	for _, pod := range summary.Pods {
		m.SetStorageUsageBytes(
			pod.PodRef.Name,
			pod.PodRef.Namespace,
			summary.Node.NodeName,
			"container", // Assuming a single container for simplicity
			float64(pod.EphemeralStorage.UsedBytes),
		)
	}
}

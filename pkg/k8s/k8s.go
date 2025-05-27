package k8s

import (
	"fmt"
)

const (
	// kubeletSummaryEndpoint is the endpoint for the kubelet summary API.
	kubeletSummaryEndpoint = "stats/summary"
)

// GetKubeletSummaryEndpoint constructs the URL for the kubelet summary endpoint.
func GetKubeletSummaryEndpoint() string {
	return fmt.Sprintf("https://localhost:10250/%s", kubeletSummaryEndpoint)
}

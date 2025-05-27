package runner

import (
	"net/http"
)

const (
	// kubeletSummaryEndpoint is the endpoint for the kubelet summary API.
	kubeletSummaryEndpoint = "https://localhost:10250/stats/summary"
)

// buildHTTPRequest creates a new HTTP request to the kubelet summary endpoint
// with the necessary authorization header.
func buildHTTPRequest() (*http.Request, error) {
	// create a new HTTP request to the kubelet summary endpoint
	req, err := http.NewRequest("GET", kubeletSummaryEndpoint, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

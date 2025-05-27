package runner

import (
	"net/http"

	"github.com/amirhnajafiz/k8sese/pkg/k8s"
)

// buildHTTPRequest creates a new HTTP request to the kubelet summary endpoint
// with the necessary authorization header.
func buildHTTPRequest() (*http.Request, error) {
	// create a new HTTP request to the kubelet summary endpoint
	req, err := http.NewRequest("GET", k8s.GetKubeletSummaryEndpoint(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

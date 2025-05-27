package runner

import (
	"fmt"
	"os"
)

const (
	// kubeletSummaryEndpoint is the endpoint for the kubelet summary API.
	kubeletSummaryEndpoint = "stats/summary"
	// kubeletAccessTokenPath is the path to the kubelet access token.
	kubeletAccessTokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
)

// GetKubeletSummaryEndpoint constructs the URL for the kubelet summary endpoint.
func GetKubeletSummaryEndpoint() string {
	return fmt.Sprintf("https://localhost:10250/%s", kubeletSummaryEndpoint)
}

// GetKubeletAccessTokenPath returns the service account token.
func GetKubeletAccessToken() (string, error) {
	tokenBytes, err := os.ReadFile(kubeletAccessTokenPath)
	if err != nil {
		return "", fmt.Errorf("failed to read kubelet access token: %w", err)
	}
	token := string(tokenBytes)

	return token, nil
}

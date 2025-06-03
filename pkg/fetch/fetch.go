package fetch

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// kubeletClientCert and key files for the kubelet client.
	certFile = "/var/lib/kubelet/pki/kubelet-client-current.pem"
	keyFile  = "/var/lib/kubelet/pki/kubelet-client-current.pem"
)

// GET performs an HTTP GET request using the provided request object.
func GET(req *http.Request) (*http.Response, error) {
	// load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load client cert/key: %w", err)
	}

	// create TLS config with InsecureSkipVerify
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// JSON decodes the JSON response from the provided HTTP response object into the given interface.
func JSON(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}

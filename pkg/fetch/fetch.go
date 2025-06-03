package fetch

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

// GET performs an HTTP GET request using the provided request object.
func GET(req *http.Request) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	http2.ConfigureTransport(tr)

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

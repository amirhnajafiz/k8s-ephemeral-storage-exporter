package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type request struct {
	Host string `json:"host"`
	Data int    `json:"data"`
}

func callback(host string, data []byte) error {
	req, err := http.NewRequest(http.MethodPost, host, bytes.NewReader(data))
	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("host=%s, code=%d %s", host, resp.StatusCode, resp.Status))

	return nil
}

func process(r *request) {}

func handler(w http.ResponseWriter, r *http.Request) {
	process(nil)
}

func main() {

}

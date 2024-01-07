package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type request struct {
	Host string `json:"host"`
	Data int    `json:"data"`
}

func callback(_ http.ResponseWriter, r *http.Request) {
	array := make([]byte, 0)

	_, _ = r.Body.Read(array)

	log.Println(string(array))
}

func handler(host string, hook string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := strconv.Atoi(r.URL.Query().Get("data"))

		body, err := json.Marshal(request{
			Host: host,
			Data: data,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		req, err := http.NewRequest(http.MethodPost, hook, bytes.NewReader(body))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		client := &http.Client{}

		if resp, err := client.Do(req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		} else if resp.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusServiceUnavailable)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	var PortFlag = flag.Int("port", 8081, "http port of app service")

	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler("", ""))
	mux.HandleFunc("/callback", callback)

	log.Println(fmt.Sprintf("app server started on %d ...", *PortFlag))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *PortFlag), mux); err != nil {
		panic(err)
	}
}

package main

import "net/http"

type request struct {
	Host   string `json:"host"`
	Method string `json:"method"`
	Data   int    `json:"data"`
}

func callback() error {
	return nil
}

func process(r *request) {}

func handler(w http.ResponseWriter, r *http.Request) {
	process(nil)
}

func main() {

}

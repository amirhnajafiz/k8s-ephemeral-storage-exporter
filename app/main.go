package main

import (
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

func callback(w http.ResponseWriter, r *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	data, _ := strconv.Atoi(r.URL.Query().Get("data"))

}

func main() {
	var PortFlag = flag.Int("port", 8081, "http port of app service")

	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)
	mux.HandleFunc("/callback", callback)

	log.Println(fmt.Sprintf("app server started on %d ...", *PortFlag))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *PortFlag), mux); err != nil {
		panic(err)
	}
}

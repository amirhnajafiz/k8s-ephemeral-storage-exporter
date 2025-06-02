package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/stats/summary", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("mock/res.json")
		if err != nil {
			http.Error(w, "could not open res.json", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte{'\n'}); err != nil {
			log.Println("error writing newline:", err)
		}
		if err := json.NewEncoder(w).Encode(json.RawMessage(mustReadAll(f))); err != nil {
			http.Error(w, "could not encode json", http.StatusInternalServerError)
		}

		log.Println("Handled /stats/summary request")
	})

	log.Println("Mock kubelet API server running on :10250")
	log.Fatal(http.ListenAndServeTLS(":10250", "mock/server.crt", "mock/server.key", nil))
}

func mustReadAll(f *os.File) []byte {
	data, err := os.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}
	return data
}

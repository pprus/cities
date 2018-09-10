package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type cityCount struct {
	Count int
	City  string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var cc cityCount

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cc); err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	log.Println(cc)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

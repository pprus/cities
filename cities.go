package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var cities = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "city_total",
		Help: "Per-city Counts",
	},
	[]string{"city"},
)

type cityCount struct {
	Count float64
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

	cities.WithLabelValues(cc.City).Add(cc.Count)
}

func init() {
	prometheus.MustRegister(cities)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

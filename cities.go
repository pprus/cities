package main

import (
	"encoding/json"
	"log"
	"net/http"
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

	log.Fatal(http.ListenAndServe(":8080", nil))
}

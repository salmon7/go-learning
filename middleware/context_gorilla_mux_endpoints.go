package main

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"log"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
/*
from https://github.com/gorilla/mux#testing-handlers
 */
package main

import (
	"log"
	"net/http"
)

// Report500 reports an Internal Server Error (HTTP_500) to the client
func Report500(w *http.ResponseWriter, message string) {
	log.Printf(message)

	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Write([]byte("500 Internal Server Error"))
}

// Report503 reports a Service Unavailable Error (HTTP_503) to the client
func Report503(w *http.ResponseWriter, message string) {
	log.Printf(message)

	(*w).WriteHeader(http.StatusServiceUnavailable)
	(*w).Write([]byte("503 Service Unavailable"))
}

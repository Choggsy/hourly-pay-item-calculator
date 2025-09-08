package web

import (
	"log"
	"net/http"
)

const server = ":8080"

// WebEntry Web server entry point
func WebEntry() {
	http.HandleFunc("/", WebHandler)
	http.ListenAndServe(server, nil)

	ErrorHandler()
	log.Println("Starting server on " + server + "...")
}

func ErrorHandler() {
	err := http.ListenAndServe(server, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

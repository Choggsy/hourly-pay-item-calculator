package main

import (
	"hourly-pay-item-calculator/web/handler"
	"log"
	"net/http"
)

const server = ":8080"

// main Web server entry point
func main() {
	http.HandleFunc("/", handler.WebHandler)
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

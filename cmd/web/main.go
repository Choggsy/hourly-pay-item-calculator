package main

import (
	"hourly-pay-item-calculator/web/handler"
	"log"
	"net/http"
)

const server = ":8080"

// main Web server entry point
func main() {
	http.HandleFunc("/", handler.WebHandler) //Registers a route, like @RequestMapping("/") in Spring Boot.
	log.Println("Starting server on " + server + "...")

	err := http.ListenAndServe(server, nil) //Starts the 8080 server
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

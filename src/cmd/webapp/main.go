package main

import (
	"log"
	"net/http"

	"cmd/webapp/infrastructure/boot"
)

// main is the entrypoint for the application.
func main() {
	// Register the services and load the routes.
	http.Handle("/", boot.ServicesAndRoutes())

	// Display message on the server.
	log.Println("Server started.")

	// Run the web listener.
	http.ListenAndServe(":8080", nil)
}

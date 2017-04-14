package main

import (
	"log"
	"net/http"
	"cmd/dbservice/infrastructure/boot"
	//"cmd/dbservice/infrastructure/boot"
)

// main is the entrypoint for the application.
func main() {
	// Register the services and load the routes.
	//boot.RegisterServices()
	handler := boot.ServicesAndRoutes()
	http.Handle("/", handler)

	// Display message on the server.
	log.Println("Server started.")

	// Run the web listener.
	http.ListenAndServe(":8081", handler)
}
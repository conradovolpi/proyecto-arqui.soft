package main

import (
	"backend/app"
	"backend/clients"
	"backend/initializers"
)

func main() {
	// Initialize the application
	initializers.LoadEnvVariables()
	clients.ConnectDb()

	// Start the router
	app.StartRoute()
}

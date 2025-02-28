package main

import (
	"log"
	"net/http"
	"golang/config"
	"golang/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	config.ConnectDB()

	// Setup routes
	r := routes.SetupRoutes()

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"log"
	"net/http"
	"golang/config"
	"golang/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	r := routes.SetupRoutes()
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}


package main

import (
	"log"
	"net/http"
	"golang/config"
	"golang/routes"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	r := routes.SetupRoutes()
	// Set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins, modify as needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

		handler := c.Handler(r)

	// fmt.Println("Server is running on port 8082")
	http.ListenAndServe(":8080", handler)
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

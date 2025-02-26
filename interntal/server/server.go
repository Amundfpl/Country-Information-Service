package server

import (
	"Assignment_1/interntal/utils"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

// Automatically load environment variables from .env
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(" Warning: No .env file found. Using system environment variables.")
	}
}

// StartServer starts the HTTP server
func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = utils.DefaultPort
	}

	r := InitializeRoutes()

	log.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err.Error())
	}
}

package server

import (
	"Assignment_1/router"
	"log"
	"net/http"
	"os"
)

// StartServer starts the HTTP server
func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	r := router.InitializeRoutes()

	log.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err.Error())
	}
}

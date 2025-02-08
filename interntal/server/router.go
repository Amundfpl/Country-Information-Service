package server

import (
	"Assignment_1/handlers"
	"net/http"
)

// InitializeRoutes sets up all API routes
func InitializeRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// ✅ Register Home Page (API Documentation)
	router.HandleFunc("/", handlers.HomeHandler)

	// Register API Endpoints
	router.HandleFunc("/countryinfo/v1/info/{countryCode}", handlers.GetCountryInfoHandler)
	router.HandleFunc("/countryinfo/v1/population/{countryCode}", handlers.GetPopulationByYearRangeHandler)
	router.HandleFunc("/countryinfo/v1/status", handlers.StatusHandler)

	return router
}

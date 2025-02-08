package router

import (
	"Assignment_1/handler"
	"net/http"
)

// InitializeRoutes sets up all API routes
func InitializeRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// ✅ Register Home Page (API Documentation)
	router.HandleFunc("/", handler.HomeHandler)

	// Register API Endpoints
	router.HandleFunc("/countryinfo/v1/info/{countryCode}", handler.GetCountryInfoHandler)
	router.HandleFunc("/countryinfo/v1/population/{countryCode}", handler.GetPopulationByYearRangeHandler)
	router.HandleFunc("/countryinfo/v1/status", handler.StatusHandler)

	return router
}

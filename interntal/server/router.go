package server

import (
	"Assignment_1/handlers"
	"Assignment_1/interntal/utils"
	"net/http"
)

// InitializeRoutes sets up all API routes
func InitializeRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Register Home Page (API Documentation)
	router.HandleFunc(utils.HomeRoute, handlers.HomeHandler)

	// Register API Endpoints
	router.HandleFunc(utils.CountryInfoRoute, handlers.GetCountryInfoHandler)
	router.HandleFunc(utils.PopulationRoute, handlers.GetPopulationByYearRangeHandler)
	router.HandleFunc(utils.StatusRoute, handlers.StatusHandler)

	return router
}

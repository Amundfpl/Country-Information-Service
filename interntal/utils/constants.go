package utils

import (
	"log"
	"os"
	"strconv"
)

// Load environment variables with fallbacks
var (
	HomeRoute        = "/"
	CountryInfoRoute = "/countryinfo/v1/info/{countryCode}"
	PopulationRoute  = "/countryinfo/v1/population/{countryCode}"
	StatusRoute      = "/countryinfo/v1/status"

	// API URLs (loaded from environment variables)
	RestCountriesAPI = getEnv("REST_COUNTRIES_API", "http://129.241.150.113:8080/v3.1")
	CountriesNowAPI  = getEnv("COUNTRIES_NOW_API", "http://129.241.150.113:3500/api/v0.1")

	// API Endpoints
	RestCountriesByAlpha   = "/alpha/"
	RestCountriesAll       = "/all"
	CountriesNowCities     = "/countries/cities"
	CountriesNowPopulation = "/countries/population"

	// Default Settings
	DefaultPort  = getEnv("PORT", "8080")
	DefaultLimit = getEnvInt("DEFAULT_LIMIT", 10)
)

// Helper function to get environment variable or fallback to default
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Helper function to get environment variable as integer
func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Warning: Invalid integer for %s. Using default: %d", key, defaultValue)
		return defaultValue
	}
	return value
}

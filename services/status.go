package services

import (
	"Assignment_1/models"
	"fmt"
	"net/http"
	"time"
)

// Store service start time
var startTime = time.Now()

// FetchServiceStatus gathers API statuses and uptime
func FetchServiceStatus() models.StatusResponse {
	// Step 1: Check API statuses
	countriesNowAPIStatus := checkAPIStatus("http://129.241.150.113:3500/api/v0.1/countries/population")
	restCountriesAPIStatus := checkAPIStatus("http://129.241.150.113:8080/v3.1/all")

	// Step 2: Calculate uptime
	uptime := time.Now().Unix() - startTime.Unix()

	// Step 3: Build structured response
	return models.StatusResponse{
		CountriesNowAPI:  countriesNowAPIStatus,
		RestCountriesAPI: restCountriesAPIStatus,
		Version:          "v1",
		Uptime:           uptime,
	}
}

// Helper function to check API status **with status code**
func checkAPIStatus(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "DOWN"
	}
	defer resp.Body.Close()

	// Returns both **status code and text** (e.g., "200 OK")
	return fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
}

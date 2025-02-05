package StatusHandler

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// Start time to track uptime
var startTime = time.Now()

// StatusResponse struct matching the expected response format
type StatusResponse struct {
	CountriesNowAPI  int    `json:"countriesnowapi"`
	RestCountriesAPI int    `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}

// Check API status by sending a HEAD request
func checkAPIStatus(url string) int {
	resp, err := http.Head(url) // Use HEAD to just get status without fetching data
	if err != nil {
		return http.StatusServiceUnavailable // Return 503 if the request fails
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

// StatusHandler handles API health checks
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// URLs of the APIs to check
	countriesNowURL := "http://129.241.150.113:3500/api/v0.1/countries"
	restCountriesURL := "http://129.241.150.113:8080/v3.1/all"

	// Use goroutines and WaitGroup to check APIs concurrently
	var wg sync.WaitGroup
	var countriesNowStatus, restCountriesStatus int

	wg.Add(2)

	// Check CountriesNow API
	go func() {
		defer wg.Done()
		countriesNowStatus = checkAPIStatus(countriesNowURL)
	}()

	// Check REST Countries API
	go func() {
		defer wg.Done()
		restCountriesStatus = checkAPIStatus(restCountriesURL)
	}()

	// Wait for both API checks to complete
	wg.Wait()

	// Calculate uptime in seconds
	uptime := int64(time.Since(startTime).Seconds())

	// Prepare response
	response := StatusResponse{
		CountriesNowAPI:  countriesNowStatus,
		RestCountriesAPI: restCountriesStatus,
		Version:          "v1",
		Uptime:           uptime,
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

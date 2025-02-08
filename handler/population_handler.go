package handler

import (
	"Assignment_1/services"
	"encoding/json"
	"net/http"
)

// GetPopulationByYearRangeHandler handles API requests for population data
func GetPopulationByYearRangeHandler(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("countryCode")
	yearRange := r.URL.Query().Get("limit") // e.g., "2010-2015"

	if countryCode == "" {
		http.Error(w, "Country code is required", http.StatusBadRequest)
		return
	}

	// Call the service function to fetch population data
	response, err := services.FetchPopulationByYearRange(countryCode, yearRange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON Response
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return
	}
	w.Write(prettyJSON)
}

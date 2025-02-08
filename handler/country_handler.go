package handler

import (
	"Assignment_1/services"
	"encoding/json"
	"net/http"
)

// GetCountryInfoHandler handles API requests for country details
func GetCountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("countryCode")
	limitStr := r.URL.Query().Get("limit") // Optional limit parameter

	if countryCode == "" {
		http.Error(w, "Country code is required", http.StatusBadRequest)
		return
	}

	// Call the service function
	countryInfo, err := services.FetchCountryInfo(countryCode, limitStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON Response
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(countryInfo, "", "    ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return
	}
	w.Write(prettyJSON)
}

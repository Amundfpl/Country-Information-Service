package handlers

import (
	"Assignment_1/interntal/utils"
	"Assignment_1/pkg/services"
	"net/http"
)

// GetCountryInfoHandler handles API requests for country details
func GetCountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("countryCode")
	limitStr := r.URL.Query().Get("limit")

	// Fetch country info from the service layer
	countryInfo, err := services.FetchCountryInfo(countryCode, limitStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Use utility function for JSON response with error checking
	if err1 := utils.RespondWithJSON(w, countryInfo); err1 != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

package handlers

import (
	"Assignment_1/interntal/utils"
	"Assignment_1/pkg/services"
	"net/http"
)

// GetPopulationByYearRangeHandler handles API requests for population data
func GetPopulationByYearRangeHandler(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("countryCode")
	yearRange := r.URL.Query().Get("limit")

	if countryCode == "" {
		http.Error(w, "Country code is required", http.StatusBadRequest)
		return
	}

	response, err := services.FetchPopulationByYearRange(countryCode, yearRange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, response) // ✅ Use helper function
}

package CountriesNowAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Handler to fetch population data by year
func GetPopulationByYearHandler(w http.ResponseWriter, r *http.Request) {
	countryName := r.PathValue("country")
	yearStr := r.PathValue("year")

	if countryName == "" || yearStr == "" {
		http.Error(w, "Country name and year are required", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid year format", http.StatusBadRequest)
		return
	}

	requestBody, _ := json.Marshal(map[string]string{"country": countryName})
	apiURL := "http://129.241.150.113:3500/api/v0.1/countries/population"
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, "Failed to fetch population data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode response
	var popData PopulationResponse
	err = json.NewDecoder(resp.Body).Decode(&popData)
	if err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	// Filter for the requested year
	var populationForYear *PopulationCounts
	for _, pop := range popData.Data.PopulationCounts {
		if pop.Year == year {
			populationForYear = &pop
			break
		}
	}

	if populationForYear == nil {
		http.Error(w, fmt.Sprintf("No population data found for %d", year), http.StatusNotFound)
		return
	}

	// Respond with filtered data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(populationForYear)
}

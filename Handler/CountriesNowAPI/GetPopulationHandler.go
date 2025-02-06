package CountriesNowAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetPopulationHandler(w http.ResponseWriter, r *http.Request) {
	countryName := r.PathValue("country")
	yearStr := r.URL.Query().Get("year") // Extract optional "year" from query params

	if countryName == "" {
		http.Error(w, "Country name is required", http.StatusBadRequest)
		return
	}

	// Handle optional "year"
	var year int
	var err error
	if yearStr != "" {
		year, err = strconv.Atoi(yearStr)
		if err != nil {
			http.Error(w, "Invalid year format", http.StatusBadRequest)
			return
		}
	}

	// API request (unchanged)
	requestBody, _ := json.Marshal(map[string]string{"country": countryName})
	apiURL := "http://129.241.150.113:3500/api/v0.1/countries/population"
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, "Failed to fetch population data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var popData PopulationResponse
	err = json.NewDecoder(resp.Body).Decode(&popData)
	if err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	// If a specific year is requested, filter the result
	if year != 0 {
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(populationForYear)
		return
	}

	// If no year is specified, return full data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(popData.Data)
}

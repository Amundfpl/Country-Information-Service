package CountriesNowAPI

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func GetPopulationHandler(w http.ResponseWriter, r *http.Request) {
	countryName := r.PathValue("country")
	if countryName == "" {
		http.Error(w, "Country name is required", http.StatusBadRequest)
		return
	}

	// Make API request
	requestBody, _ := json.Marshal(map[string]string{"country": countryName})
	apiURL := "http://129.241.150.113:3500/api/v0.1/countries/population"
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, "Failed to fetch population data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode JSON response
	var popData PopulationResponse
	err = json.NewDecoder(resp.Body).Decode(&popData)
	if err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	// Return the full population data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(popData.Data)
}

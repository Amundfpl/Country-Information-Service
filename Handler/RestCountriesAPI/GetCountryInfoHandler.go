package RestCountriesAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define the struct based on the API response
type CountryInfo struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Population int               `json:"population"`
	Region     string            `json:"region"`
	Subregion  string            `json:"subregion"`
	Languages  map[string]string `json:"languages"`
}

func GetCountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	countryName := r.PathValue("country")
	if countryName == "" {
		http.Error(w, "Country name is required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Fetching country info for %s", countryName)

	apiURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/name/%s", countryName)
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch country info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Country not found", resp.StatusCode)
		return
	}

	var countryData []CountryInfo
	err = json.NewDecoder(resp.Body).Decode(&countryData)
	if err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countryData[0])
}

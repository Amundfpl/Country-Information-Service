package services

import (
	"Assignment_1/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

// FetchCountryInfo retrieves country data using country code
func FetchCountryInfo(countryCode string, limit int) (*models.CountryInfo, error) {
	// Fetch country details
	restCountriesURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(restCountriesURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch country info: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var countryData []models.CountryInfo
	err = json.Unmarshal(bodyBytes, &countryData)
	if err != nil || len(countryData) == 0 {
		return nil, fmt.Errorf("failed to decode country info")
	}

	fullCountryName := countryData[0].Name

	// Fetch cities from CountriesNow API
	citiesAPIURL := "http://129.241.150.113:3500/api/v0.1/countries/cities"
	requestBody, _ := json.Marshal(map[string]string{"country": fullCountryName})

	citiesResp, err := http.Post(citiesAPIURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cities")
	}
	defer citiesResp.Body.Close()

	citiesBytes, _ := io.ReadAll(citiesResp.Body)

	var citiesData models.CitiesResponse
	err = json.Unmarshal(citiesBytes, &citiesData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode cities")
	}

	// Sort and apply limit
	sort.Strings(citiesData.Data)
	if len(citiesData.Data) > limit {
		citiesData.Data = citiesData.Data[:limit]
	}

	countryData[0].Cities = citiesData.Data
	return &countryData[0], nil
}

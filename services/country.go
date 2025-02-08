package services

import (
	"Assignment_1/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
)

// FetchCountryInfo retrieves country details and city list based on country code
func FetchCountryInfo(countryCode, limitStr string) (models.CountryInfo, error) {
	limit, err := parseLimit(limitStr, 10) // Default limit: 10
	if err != nil {
		return models.CountryInfo{}, err
	}

	fullCountryName, countryData, err := getCountryData(countryCode)
	if err != nil {
		return models.CountryInfo{}, err
	}

	cities, err := fetchCities(fullCountryName, limit)
	if err != nil {
		return models.CountryInfo{}, err
	}

	return formatCountryResponse(fullCountryName, countryData, cities), nil
}

// parseLimit validates and converts the limit parameter
func parseLimit(limitStr string, defaultLimit int) (int, error) {
	if limitStr == "" {
		return defaultLimit, nil
	}

	parsedLimit, err := strconv.Atoi(limitStr)
	if err != nil || parsedLimit <= 0 {
		return 0, fmt.Errorf("invalid limit parameter")
	}

	return parsedLimit, nil
}

// getCountryData fetches country details from REST Countries API
func getCountryData(countryCode string) (string, models.CountryInfoResponse, error) {
	url := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", models.CountryInfoResponse{}, fmt.Errorf("failed to fetch country info")
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var countryData []models.CountryInfoResponse
	if err := json.Unmarshal(bodyBytes, &countryData); err != nil || len(countryData) == 0 {
		return "", models.CountryInfoResponse{}, fmt.Errorf("failed to decode country info")
	}

	return countryData[0].Name.Common, countryData[0], nil
}

// fetchCities retrieves and sorts city names from CountriesNow API
func fetchCities(countryName string, limit int) ([]string, error) {
	apiURL := "http://129.241.150.113:3500/api/v0.1/countries/cities"
	requestBody, _ := json.Marshal(map[string]string{"country": countryName})

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch cities")
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var citiesData models.CitiesResponse
	if err := json.Unmarshal(bodyBytes, &citiesData); err != nil {
		return nil, fmt.Errorf("failed to decode cities")
	}

	sort.Strings(citiesData.Data) // Alphabetical sorting

	if len(citiesData.Data) > limit {
		citiesData.Data = citiesData.Data[:limit] // Apply limit
	}

	return citiesData.Data, nil
}

// formatCountryResponse structures the final API response
func formatCountryResponse(countryName string, countryData models.CountryInfoResponse, cities []string) models.CountryInfo {
	capital := ""
	if len(countryData.Capital) > 0 {
		capital = countryData.Capital[0]
	}

	return models.CountryInfo{
		Name:       countryName,
		Continents: countryData.Continents,
		Population: countryData.Population,
		Languages:  countryData.Languages,
		Borders:    countryData.Borders,
		Flag:       countryData.Flags.Png,
		Capital:    capital,
		Cities:     cities,
	}
}

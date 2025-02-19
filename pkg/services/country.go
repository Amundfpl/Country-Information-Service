package services

import (
	"Assignment_1/interntal/utils"
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
	limit, err := parseLimit(limitStr, utils.DefaultLimit) // Default limit: 10
	if err != nil {
		return models.CountryInfo{}, err
	}

	fullCountryName, countryData, err1 := getCountryData(countryCode)
	if err1 != nil {
		return models.CountryInfo{}, err1
	}

	cities, err2 := fetchCities(fullCountryName, limit)
	if err2 != nil {
		return models.CountryInfo{}, err2
	}

	return formatCountryResponse(fullCountryName, countryData, cities), nil
}

// parseLimit validates and converts the limit parameter
func parseLimit(limitStr string, defaultLimit int) (int, error) {
	if limitStr == "" {
		return defaultLimit, nil
	}

	parsedLimit, err3 := strconv.Atoi(limitStr)
	if err3 != nil || parsedLimit <= 0 {
		return 0, fmt.Errorf("invalid limit parameter")
	}

	return parsedLimit, nil
}

// getCountryData fetches country details from REST Countries API
func getCountryData(countryCode string) (string, models.CountryInfoResponse, error) {
	url := fmt.Sprintf("%s%s%s", utils.RestCountriesAPI, utils.RestCountriesByAlpha, countryCode)
	resp, err4 := http.Get(url)
	if err4 != nil || resp.StatusCode != http.StatusOK {
		return "", models.CountryInfoResponse{}, fmt.Errorf("failed to fetch country info")
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var countryData []models.CountryInfoResponse
	if err5 := json.Unmarshal(bodyBytes, &countryData); err5 != nil || len(countryData) == 0 {
		return "", models.CountryInfoResponse{}, fmt.Errorf("failed to decode country info")
	}

	return countryData[0].Name.Common, countryData[0], nil
}

// fetchCities retrieves and sorts city names from CountriesNow API
func fetchCities(countryName string, limit int) ([]string, error) {
	apiURL := fmt.Sprintf("%s%s", utils.CountriesNowAPI, utils.CountriesNowCities)
	requestBody, _ := json.Marshal(map[string]string{"country": countryName})

	resp, err6 := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err6 != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch cities")
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var citiesData models.CitiesResponse
	if err7 := json.Unmarshal(bodyBytes, &citiesData); err7 != nil {
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

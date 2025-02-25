package services

import (
	"Assignment_1/interntal/utils"
	"Assignment_1/models"
	"fmt"
	"sort"
	"strconv"
)

// FetchCountryInfo retrieves country details and city list based on country code
func FetchCountryInfo(countryCode, limitStr string) (models.CountryInfo, error) {
	limit, errLimit := parseLimit(limitStr, utils.DefaultLimit) // Default limit: 10
	if errLimit != nil {
		return models.CountryInfo{}, errLimit
	}

	fullCountryName, countryData, errCountryData := getCountryData(countryCode)
	if errCountryData != nil {
		return models.CountryInfo{}, errCountryData
	}

	cities, errCities := fetchCities(fullCountryName, limit)
	if errCities != nil {
		return models.CountryInfo{}, errCities
	}

	return formatCountryResponse(fullCountryName, countryData, cities), nil
}

// parseLimit validates and converts the limit parameter
func parseLimit(limitStr string, defaultLimit int) (int, error) {
	if limitStr == "" {
		return defaultLimit, nil
	}

	parsedLimit, errParse := strconv.Atoi(limitStr)
	if errParse != nil || parsedLimit <= 0 {
		return 0, fmt.Errorf("invalid limit parameter: %v", errParse)
	}

	return parsedLimit, nil
}

// getCountryData fetches country details from REST Countries API
func getCountryData(countryCode string) (string, models.CountryInfoResponse, error) {
	url := fmt.Sprintf("%s%s%s", utils.RestCountriesAPI, utils.RestCountriesByAlpha, countryCode)

	var countryData []models.CountryInfoResponse
	err := utils.GetRequest(url, &countryData)
	if err != nil || len(countryData) == 0 {
		return "", models.CountryInfoResponse{}, fmt.Errorf("failed to fetch country info: %v", err)
	}

	return countryData[0].Name.Common, countryData[0], nil
}

// fetchCities retrieves and sorts city names from CountriesNow API
func fetchCities(countryName string, limit int) ([]string, error) {
	apiURL := fmt.Sprintf("%s%s", utils.CountriesNowAPI, utils.CountriesNowCities)

	var citiesData models.CitiesResponse
	err := utils.PostRequest(apiURL, map[string]string{"country": countryName}, &citiesData)
	if err != nil {
		return nil, err
	}

	// Sort cities alphabetically
	sort.Strings(citiesData.Data)

	// Apply limit to cities
	if len(citiesData.Data) > limit {
		citiesData.Data = citiesData.Data[:limit]
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

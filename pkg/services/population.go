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
	"strings"
)

// FetchPopulationByYearRange retrieves population data within a given year range
func FetchPopulationByYearRange(countryCode, yearRange string) (map[string]interface{}, error) {
	countryCode = strings.ToUpper(countryCode)

	fullCountryName, err := getCountryName(countryCode)
	if err != nil {
		return nil, err
	}

	populationCounts, err := fetchPopulationData(fullCountryName)
	if err != nil {
		return nil, err
	}

	filteredPopulations := filterPopulationByYearRange(populationCounts, yearRange)

	meanPopulation := calculateMeanPopulation(filteredPopulations)

	return formatPopulationResponse(fullCountryName, filteredPopulations, meanPopulation), nil
}

// getCountryName fetches the full country name from REST Countries API
func getCountryName(countryCode string) (string, error) {
	url := fmt.Sprintf("%s%s%s", utils.RestCountriesAPI, utils.RestCountriesByAlpha, countryCode)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch country name: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var countryData []models.CountryInfoResponse
	if err := json.Unmarshal(bodyBytes, &countryData); err != nil || len(countryData) == 0 {
		return "", fmt.Errorf("failed to decode country info")
	}

	return countryData[0].Name.Common, nil
}

// fetchPopulationData gets population data from CountriesNow API
func fetchPopulationData(countryName string) ([]models.PopulationCounts, error) {
	apiURL := fmt.Sprintf("%s%s", utils.CountriesNowAPI, utils.CountriesNowPopulation)
	requestBody, _ := json.Marshal(map[string]string{"country": countryName})

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch population data")
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var popData models.PopulationResponse
	if err := json.Unmarshal(bodyBytes, &popData); err != nil || popData.Error {
		return nil, fmt.Errorf("population data not found")
	}

	sort.Slice(popData.Data.PopulationCounts, func(i, j int) bool {
		return popData.Data.PopulationCounts[i].Year < popData.Data.PopulationCounts[j].Year
	})

	return popData.Data.PopulationCounts, nil
}

// filterPopulationByYearRange filters population data within a given year range
func filterPopulationByYearRange(populationCounts []models.PopulationCounts, yearRange string) []models.PopulationCounts {
	if yearRange == "" {
		return populationCounts
	}

	yearParts := strings.Split(yearRange, "-")
	if len(yearParts) != 2 {
		return nil
	}

	startYear, _ := strconv.Atoi(yearParts[0])
	endYear, _ := strconv.Atoi(yearParts[1])

	if startYear > endYear || startYear == 0 || endYear == 0 {
		return nil
	}

	filtered := []models.PopulationCounts{}
	for _, pop := range populationCounts {
		if pop.Year >= startYear && pop.Year <= endYear {
			filtered = append(filtered, pop)
		}
	}

	return filtered
}

// calculateMeanPopulation computes the mean population over the filtered years
func calculateMeanPopulation(populations []models.PopulationCounts) int {
	if len(populations) == 0 {
		return 0
	}

	totalPopulation := 0
	for _, pop := range populations {
		totalPopulation += pop.Value
	}

	return totalPopulation / len(populations)
}

// formatPopulationResponse structures the final API response
func formatPopulationResponse(countryName string, populations []models.PopulationCounts, meanPopulation int) map[string]interface{} {
	return map[string]interface{}{
		"country":        countryName,
		"populationData": populations,
		"meanPopulation": meanPopulation,
	}
}

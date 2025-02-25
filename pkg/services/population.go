package services

import (
	"Assignment_1/interntal/utils"
	"Assignment_1/models"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// FetchPopulationByYearRange retrieves population data within a given year range
func FetchPopulationByYearRange(countryCode, yearRange string) (map[string]interface{}, error) {
	countryCode = strings.ToUpper(countryCode)

	fullCountryName, errCountryName := getCountryName(countryCode)
	if errCountryName != nil {
		return nil, errCountryName
	}

	populationCounts, errPopulationData := fetchPopulationData(fullCountryName)
	if errPopulationData != nil {
		return nil, errPopulationData
	}

	filteredPopulations := filterPopulationByYearRange(populationCounts, yearRange)

	meanPopulation := calculateMeanPopulation(filteredPopulations)

	return formatPopulationResponse(fullCountryName, filteredPopulations, meanPopulation), nil
}

// getCountryName fetches the full country name from REST Countries API
func getCountryName(countryCode string) (string, error) {
	url := fmt.Sprintf("%s%s%s", utils.RestCountriesAPI, utils.RestCountriesByAlpha, countryCode)

	var countryData []models.CountryInfoResponse
	err := utils.GetRequest(url, &countryData)
	if err != nil || len(countryData) == 0 {
		return "", fmt.Errorf("failed to fetch country info: %v", err)
	}

	return countryData[0].Name.Common, nil
}

// fetchPopulationData gets population data from CountriesNow API
func fetchPopulationData(countryName string) ([]models.PopulationCounts, error) {
	apiURL := fmt.Sprintf("%s%s", utils.CountriesNowAPI, utils.CountriesNowPopulation)

	var popData models.PopulationResponse
	err := utils.PostRequest(apiURL, map[string]string{"country": countryName}, &popData)
	if err != nil {
		return nil, err
	}

	if popData.Error {
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

	startYear, errStartYear := strconv.Atoi(yearParts[0])
	if errStartYear != nil {
		fmt.Println("Warning: invalid start year format:", errStartYear)
		return nil
	}

	endYear, errEndYear := strconv.Atoi(yearParts[1])
	if errEndYear != nil {
		fmt.Println("Warning: invalid end year format:", errEndYear)
		return nil
	}

	if startYear > endYear || startYear == 0 || endYear == 0 {
		fmt.Println("Warning: invalid year range:", yearRange)
		return nil
	}

	var filtered []models.PopulationCounts
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

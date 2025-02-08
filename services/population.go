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
	"strings"
)

// FetchPopulationByYearRange retrieves population data within a given year range
func FetchPopulationByYearRange(countryCode, yearRange string) (map[string]interface{}, error) {
	// Convert countryCode to uppercase
	countryCode = strings.ToUpper(countryCode)

	// Fetch country name from REST Countries API
	restCountriesURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(restCountriesURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch country name: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var countryData []models.CountryInfoResponse
	err = json.Unmarshal(bodyBytes, &countryData)
	if err != nil || len(countryData) == 0 {
		return nil, fmt.Errorf("failed to decode country info")
	}

	fullCountryName := countryData[0].Name.Common // Extract full country name

	// Fetch Population Data from CountriesNow API
	apiURL := "http://129.241.150.113:3500/api/v0.1/countries/population"
	requestBody, _ := json.Marshal(map[string]string{"country": fullCountryName})

	popResp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch population data")
	}
	defer popResp.Body.Close()

	popBodyBytes, _ := io.ReadAll(popResp.Body)

	var popData models.PopulationResponse
	err = json.Unmarshal(popBodyBytes, &popData)
	if err != nil || popData.Error {
		return nil, fmt.Errorf("population data not found")
	}

	// Process Population Data
	populationCounts := popData.Data.PopulationCounts
	sort.Slice(populationCounts, func(i, j int) bool {
		return populationCounts[i].Year < populationCounts[j].Year
	})

	// Filter Population Data Based on Year Range
	filteredPopulations := populationCounts
	startYear, endYear := 0, 0
	if yearRange != "" {
		yearParts := strings.Split(yearRange, "-")
		if len(yearParts) == 2 {
			startYear, _ = strconv.Atoi(yearParts[0])
			endYear, _ = strconv.Atoi(yearParts[1])

			if startYear > endYear || startYear == 0 || endYear == 0 {
				return nil, fmt.Errorf("invalid year range")
			}

			// Filter population data
			filteredPopulations = []models.PopulationCounts{}
			for _, pop := range populationCounts {
				if pop.Year >= startYear && pop.Year <= endYear {
					filteredPopulations = append(filteredPopulations, pop)
				}
			}
		}
	}

	// Calculate Mean Population
	totalPopulation := 0
	count := len(filteredPopulations)
	if count > 0 {
		for _, pop := range filteredPopulations {
			totalPopulation += pop.Value
		}
	}

	meanPopulation := 0
	if count > 0 {
		meanPopulation = totalPopulation / count
	}

	// Return processed data
	response := map[string]interface{}{
		"country":        fullCountryName,
		"populationData": filteredPopulations,
		"meanPopulation": meanPopulation,
	}

	return response, nil
}

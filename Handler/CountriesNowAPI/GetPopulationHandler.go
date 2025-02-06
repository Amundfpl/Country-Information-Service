package CountriesNowAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Struct for REST Countries API response
type CountryInfoResponse struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
}

// Struct for CountriesNow API response

func GetPopulationByYearRangeHandler(w http.ResponseWriter, r *http.Request) {
	countryCode := strings.ToUpper(r.PathValue("countryCode")) // Ensure uppercase (e.g., "NO")
	yearRange := r.URL.Query().Get("limit")                    // e.g., "2010-2015"

	if countryCode == "" {
		http.Error(w, "Country code is required", http.StatusBadRequest)
		return
	}

	// ✅ Step 1: Get Full Country Name from REST Countries API
	restCountriesURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(restCountriesURL)
	if err != nil {
		http.Error(w, "Failed to fetch country name", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var countryData []CountryInfoResponse
	err = json.Unmarshal(bodyBytes, &countryData)
	if err != nil || len(countryData) == 0 {
		http.Error(w, "Failed to decode country info", http.StatusInternalServerError)
		return
	}

	fullCountryName := countryData[0].Name.Common          // Extract full country name
	fmt.Println("Resolved Country Name:", fullCountryName) // Debugging

	// ✅ Step 2: Fetch Population Data from CountriesNow API using full country name
	apiURL := "http://129.241.150.113:3500/api/v0.1/countries/population"
	requestBody, _ := json.Marshal(map[string]string{"country": fullCountryName})

	popResp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, "Failed to fetch population data", http.StatusInternalServerError)
		return
	}
	defer popResp.Body.Close()

	popBodyBytes, _ := io.ReadAll(popResp.Body)
	fmt.Println("Raw API Response:", string(popBodyBytes)) // Debugging

	var popData PopulationResponse
	err = json.Unmarshal(popBodyBytes, &popData)
	if err != nil || popData.Error {
		http.Error(w, "Population data not found", http.StatusNotFound)
		return
	}

	// ✅ Step 3: Process Population Data
	populationCounts := popData.Data.PopulationCounts
	sort.Slice(populationCounts, func(i, j int) bool {
		return populationCounts[i].Year < populationCounts[j].Year
	})

	// ✅ Step 4: Filter Population Data Based on Year Range
	filteredPopulations := populationCounts
	startYear, endYear := 0, 0
	if yearRange != "" {
		yearParts := strings.Split(yearRange, "-")
		if len(yearParts) == 2 {
			startYear, _ = strconv.Atoi(yearParts[0])
			endYear, _ = strconv.Atoi(yearParts[1])

			if startYear > endYear || startYear == 0 || endYear == 0 {
				http.Error(w, "Invalid year range", http.StatusBadRequest)
				return
			}

			// Filter population data
			filteredPopulations = []PopulationCounts{}
			for _, pop := range populationCounts {
				if pop.Year >= startYear && pop.Year <= endYear {
					filteredPopulations = append(filteredPopulations, pop)
				}
			}
		}
	}

	// ✅ Step 5: Calculate Mean Population
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

	// ✅ Step 6: Return Response
	response := map[string]interface{}{
		"country":        fullCountryName,
		"populationData": filteredPopulations,
		"meanPopulation": meanPopulation,
	}

	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return
	}
	w.Write(prettyJSON)
}

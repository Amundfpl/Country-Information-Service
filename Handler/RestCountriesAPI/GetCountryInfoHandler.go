package RestCountriesAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
)

// API struct for decoding response
type APIResponse struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       struct {
		Png string `json:"png"`
	} `json:"flags"`
	Capital []string `json:"capital"`
}

// Struct for Cities API response
type CitiesResponse struct {
	Error bool     `json:"error"`
	Msg   string   `json:"msg"`
	Data  []string `json:"data"`
}

func GetCountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("countryCode")
	limitStr := r.URL.Query().Get("limit")

	if countryCode == "" {
		http.Error(w, "Country code is required", http.StatusBadRequest)
		return
	}

	// Set default limit if not provided
	limit := 10
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			http.Error(w, "Invalid limit value", http.StatusBadRequest)
			return
		}
	}

	// ✅ Fetch country info from REST Countries API
	restCountriesURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(restCountriesURL)
	if err != nil {
		http.Error(w, "Failed to fetch country info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var countryData []APIResponse
	err = json.Unmarshal(bodyBytes, &countryData)
	if err != nil {
		http.Error(w, "Failed to decode country info", http.StatusInternalServerError)
		return
	}

	if len(countryData) == 0 {
		http.Error(w, "No country data found", http.StatusNotFound)
		return
	}

	country := countryData[0] // Take the first country

	// Extract first capital if available
	capital := ""
	if len(country.Capital) > 0 {
		capital = country.Capital[0]
	}

	// ✅ Fetch cities from CountriesNow API
	citiesAPIURL := "http://129.241.150.113:3500/api/v0.1/countries/cities"
	requestBody, _ := json.Marshal(map[string]string{"country": country.Name.Common})

	citiesResp, err := http.Post(citiesAPIURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, "Failed to fetch cities", http.StatusInternalServerError)
		return
	}
	defer citiesResp.Body.Close()

	citiesBytes, _ := io.ReadAll(citiesResp.Body)

	var citiesData CitiesResponse
	err = json.Unmarshal(citiesBytes, &citiesData)
	if err != nil {
		http.Error(w, "Failed to decode cities", http.StatusInternalServerError)
		return
	}

	cities := citiesData.Data
	sort.Strings(cities) // Sort alphabetically

	if len(cities) > limit {
		cities = cities[:limit]
	}

	// ✅ Convert to final struct
	finalResponse := map[string]interface{}{
		"name":       country.Name.Common,
		"continents": country.Continents,
		"population": country.Population,
		"languages":  country.Languages,
		"borders":    country.Borders,
		"flag":       country.Flag.Png,
		"capital":    capital,
		"cities":     cities,
	}

	// ✅ Return response with pretty JSON formatting
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(finalResponse, "", "    ") // 4-space indentation
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return
	}
	w.Write(prettyJSON)
}

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

func GetCountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	countryCode := r.PathValue("countryCode")
	limitStr := r.URL.Query().Get("limit")
	defaultLimit := 10 // Default limit if not provided

	// Convert limit to int (if provided), else use default
	limit := defaultLimit
	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	if countryCode == "" {
		http.Error(w, "Country code is required", http.StatusBadRequest)
		return
	}

	// ✅ Step 1: Fetch Country Data from REST Countries API
	restCountriesURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", countryCode)
	resp, err := http.Get(restCountriesURL)
	if err != nil {
		http.Error(w, "Failed to fetch country info", http.StatusInternalServerError)
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

	// Extract full country name
	fullCountryName := countryData[0].Name.Common

	// ✅ Step 2: Fetch Cities from CountriesNow API
	citiesAPIURL := "http://129.241.150.113:3500/api/v0.1/countries/cities"
	requestBody, _ := json.Marshal(map[string]string{"country": fullCountryName})

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

	// ✅ Step 3: Sort Cities Alphabetically
	sort.Strings(citiesData.Data)

	// ✅ Step 4: Apply Limit (Default: 10 Cities)
	if len(citiesData.Data) > limit {
		citiesData.Data = citiesData.Data[:limit]
	}

	// Extract capital safely
	capital := ""
	if len(countryData[0].Capital) > 0 {
		capital = countryData[0].Capital[0]
	}

	// ✅ Step 5: Create Struct Response
	countryInfo := CountryInfo{
		Name:       fullCountryName,
		Continents: countryData[0].Continents,
		Population: countryData[0].Population,
		Languages:  countryData[0].Languages,
		Borders:    countryData[0].Borders,
		Flag:       countryData[0].Flags.Png,
		Capital:    capital,
		Cities:     citiesData.Data,
	}

	// ✅ Step 6: Return JSON Response
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(countryInfo, "", "    ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return
	}
	w.Write(prettyJSON)
}

package utils

// API Route Paths
const (
	HomeRoute        = "/"
	CountryInfoRoute = "/countryinfo/v1/info/{countryCode}"
	PopulationRoute  = "/countryinfo/v1/population/{countryCode}"
	StatusRoute      = "/countryinfo/v1/status"
)

// Base API URLs
const (
	RestCountriesAPI = "http://129.241.150.113:8080/v3.1"     // Base URL for REST Countries API
	CountriesNowAPI  = "http://129.241.150.113:3500/api/v0.1" // Base URL for CountriesNow API
)

// REST Countries API Endpoints
const (
	RestCountriesByAlpha = "/alpha/" // Fetch country by Alpha-2/Alpha-3 code
	RestCountriesAll     = "/all"    // Fetch all countries (if needed)
)

// CountriesNow API Endpoints
const (
	CountriesNowCities     = "/countries/cities"     // Fetch cities
	CountriesNowPopulation = "/countries/population" // Fetch population data
)

// Default Settings
const (
	DefaultPort  = "8080"
	DefaultLimit = 10
)

package CountriesNowAPI

// Struct for Population API response
type PopulationResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		Country          string             `json:"country"`
		Code             string             `json:"code"`
		Iso3             string             `json:"iso3"`
		PopulationCounts []PopulationCounts `json:"populationCounts"`
	} `json:"data"`
}

type PopulationCounts struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

package services

import (
	"Assignment_1/interntal/utils"
	"Assignment_1/models"
	"fmt"
	"net/http"
	"time"
)

// Store service start time
var startTime = time.Now()

// FetchServiceStatus gathers API statuses and uptime
func FetchServiceStatus() models.StatusResponse {
	return models.StatusResponse{
		CountriesNowAPI:  getAPIStatus(utils.CountriesNowAPI + utils.CountriesNowPopulation),
		RestCountriesAPI: getAPIStatus(utils.RestCountriesAPI + utils.RestCountriesAll),
		Version:          "v1",
		Uptime:           time.Now().Unix() - startTime.Unix(),
	}
}

// getAPIStatus checks API availability and returns its status code & text
func getAPIStatus(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "DOWN"
	}
	defer resp.Body.Close()

	return fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
}

package main

import (
	"Assignment_1/Handler/CountriesNowAPI"
	"Assignment_1/Handler/RestCountriesAPI"
	"Assignment_1/Handler/StatusHandler"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	router := http.NewServeMux()

	// Corrected route definitions
	router.HandleFunc("/countryinfo/v1/info/{country}", RestCountriesAPI.GetCountryInfoHandler)
	router.HandleFunc("/countryinfo/v1/population/{country}", CountriesNowAPI.GetPopulationHandler)
	router.HandleFunc("/countryinfo/v1/population/{country}/{year}", CountriesNowAPI.GetPopulationByYearHandler)
	router.HandleFunc("/countryinfo/v1/status", StatusHandler.StatusHandler)

	log.Println("Running on port", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err.Error())
	}
}

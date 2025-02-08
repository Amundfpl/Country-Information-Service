package handlers

import (
	"Assignment_1/interntal/utils"
	"Assignment_1/pkg/services"
	"net/http"
)

// StatusHandler handles API status requests
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := services.FetchServiceStatus()
	utils.RespondWithJSON(w, status) //  Use helper function
}

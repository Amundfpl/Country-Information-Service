package handlers

import (
	"Assignment_1/interntal/utils"
	"Assignment_1/pkg/services"
	"net/http"
)

// StatusHandler handles API status requests
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := services.FetchServiceStatus()

	if err := utils.RespondWithJSON(w, status); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	} //  Use helper function
}

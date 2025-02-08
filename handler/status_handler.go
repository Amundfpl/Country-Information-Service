package handler

import (
	"Assignment_1/services"
	"encoding/json"
	"net/http"
)

// StatusHandler handles requests for API status
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch service status
	status := services.FetchServiceStatus()

	// Return JSON Response
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(status, "", "    ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return
	}
	w.Write(prettyJSON)
}

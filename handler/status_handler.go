package handler

import (
	"Assignment_1/services"
	"encoding/json"
	"net/http"
)

// StatusHandler handles the `/status` request
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Call the service function to fetch API statuses
	response := services.FetchServiceStatus()

	// Return pretty-printed JSON response
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return
	}
	w.Write(prettyJSON)
}

package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON formats and sends a properly formatted JSON response
func RespondWithJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return
	}
	w.Write(prettyJSON)
}

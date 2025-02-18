package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON formats and sends a properly formatted JSON response
func RespondWithJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return err
	}

	// Handle the error from w.Write
	if _, err := w.Write(prettyJSON); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return err
	}

	return nil
}

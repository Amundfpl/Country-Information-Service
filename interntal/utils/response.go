package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// RespondWithJSON formats and sends a properly formatted JSON response
func RespondWithJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(w, "Failed to format JSON", http.StatusInternalServerError)
		return fmt.Errorf("failed to format JSON response: %v", err)
	}

	// Handle the error from w.Write
	if _, errWrite := w.Write(prettyJSON); errWrite != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return fmt.Errorf("failed to write response: %v", errWrite)
	}

	return nil
}

// PostRequest makes a POST request and decodes the response into the provided struct.
func PostRequest(url string, body map[string]string, responseStruct interface{}) error {
	// Marshal the request body
	requestBody, errMarshal := json.Marshal(body)
	if errMarshal != nil {
		return fmt.Errorf("failed to marshal request body: %v", errMarshal)
	}

	// Send the POST request
	resp, errPost := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if errPost != nil {
		return fmt.Errorf("HTTP POST request failed: %v", errPost)
	}
	defer CloseResponseBody(resp)

	// Check if the status code is OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP POST returned status %d", resp.StatusCode)
	}

	// Decode JSON response into the provided struct
	bodyBytes, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return fmt.Errorf("failed to read response body: %v", errRead)
	}

	if errUnmarshal := json.Unmarshal(bodyBytes, responseStruct); errUnmarshal != nil {
		return fmt.Errorf("failed to decode response JSON: %v", errUnmarshal)
	}

	return nil
}

// GetRequest makes a GET request and decodes the response into the provided struct.
func GetRequest(url string, responseStruct interface{}) error {
	resp, errGet := http.Get(url)
	if errGet != nil {
		return fmt.Errorf("HTTP GET request failed: %v", errGet)
	}
	defer CloseResponseBody(resp)

	// Ensure HTTP response is OK before reading the body
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP GET to %s returned status %d", url, resp.StatusCode)
	}

	// Read response body
	bodyBytes, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return fmt.Errorf("failed to read response body: %v", errRead)
	}

	// Decode JSON response into provided struct
	if errUnmarshal := json.Unmarshal(bodyBytes, responseStruct); errUnmarshal != nil {
		return fmt.Errorf("failed to decode JSON response: %v", errUnmarshal)
	}

	return nil
}

// CloseResponseBody safely closes an HTTP response body and logs any errors.
func CloseResponseBody(resp *http.Response) {
	if resp != nil {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("Warning: Failed to close response body:", err)
		}
	}
}

// GetStatusCode makes a GET request and returns only the HTTP status code.
func GetStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("HTTP GET request failed: %v", err)
	}
	defer resp.Body.Close()

	// Ensure the response is valid
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("HTTP GET returned status %d", resp.StatusCode)
	}

	return resp.StatusCode, nil
}

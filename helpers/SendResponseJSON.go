package helpers

import (
	"encoding/json"
	"net/http"
)

// Response struct
type Response struct {
	Message string `json:"message"`
}

// SendResponseJSON takes string, encodes as JSON, returns
func SendResponseJSON(message string, w http.ResponseWriter, statusCode int) {
	response := Response{message}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

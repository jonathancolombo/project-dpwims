package utilities

import (
	"encoding/json"
	"net/http"
)

// DecodeJSON decodes the JSON body of an HTTP request into a value of type T.
func DecodeJSON[T any](writer http.ResponseWriter, request *http.Request) (*T, bool) {
	var body T
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		http.Error(writer, "invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return nil, false
	}
	return &body, true
}

// WriteJSON writes a JSON response with the given status code and body.
func WriteJSON(writer http.ResponseWriter, status int, body any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(body)
	if err != nil {
		return
	}
}

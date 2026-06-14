package utilities

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

// ParseIDParam extracts and parses an integer ID from a chi URL parameter.
func ParseIDParam(writer http.ResponseWriter, request *http.Request, param string) (int64, bool) {
	idStr := chi.URLParam(request, param)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(writer, "invalid ID: "+err.Error(), http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

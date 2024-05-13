package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func ResponseJSON(w http.ResponseWriter, statusCode int, payload interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(payload)
}

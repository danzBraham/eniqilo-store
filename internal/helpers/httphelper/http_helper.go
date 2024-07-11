package httphelper

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func DecodeJSON(r *http.Request, payload any) error {
	return json.NewDecoder(r.Body).Decode(payload)
}

func EncodeJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func ErrorResponse(w http.ResponseWriter, status int, err error) {
	EncodeJSON(w, status, ResponseBody{
		Error:   http.StatusText(status),
		Message: err.Error(),
	})
}

func SuccessResponse(w http.ResponseWriter, status int, message string, data any) {
	EncodeJSON(w, status, ResponseBody{
		Message: message,
		Data:    data,
	})
}

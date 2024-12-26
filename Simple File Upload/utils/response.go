package utils

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, status bool, message string, code int, data interface{}) {
	response := map[string]interface{}{
		"status": status,
		"message": message,
		"code": code,
		"data": data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
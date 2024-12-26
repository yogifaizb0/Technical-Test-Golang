package utils

import (
	"basicrestapi/dto"
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, status bool, message string, code int, data interface{}) {
	response := dto.BaseResponse{
		Status: status,
		Message: message,
		Code: code,
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
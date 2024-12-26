package utils

import (
	"encoding/json"
	"net/http"
	"userauthentication/dto"
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
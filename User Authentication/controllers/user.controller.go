package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"userauthentication/dto"
	"userauthentication/services"
	"userauthentication/utils"

	"github.com/golang-jwt/jwt/v5"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		utils.Response(w, false, "Invalid input", http.StatusBadRequest, nil)
		return
	}

	err := services.Register(userDto)
	if err != nil{
		utils.Response(w, false, "Failed to hash password", http.StatusInternalServerError, nil)
	}
	utils.Response(w, true, "Register success", http.StatusCreated, nil)
}

func LoginController(w http.ResponseWriter, r * http.Request) {
	var userDto dto.UserDto
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		utils.Response(w, false, "Invalid input", http.StatusBadRequest, nil)
		return
	}

	user, err := services.Login(userDto)
	if err != nil {
		utils.Response(w, false, err.Error(), http.StatusUnauthorized, nil)
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		utils.Response(w, false, err.Error(), http.StatusInternalServerError, nil)
		return
	}
	w.Header().Set("Authorization", "Bearer " + token)
	utils.Response(w, true, "Login success", http.StatusOK, nil)
}

func ProfileController(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if  authHeader == "" {
		utils.Response(w, false, "You must login first", http.StatusUnauthorized, nil)
		return
	}

	parts := strings.Fields(authHeader)
	if len(parts) != 2 || parts[0] != "Bearer" {
		utils.Response(w, false, "Invalid token format", http.StatusUnauthorized, nil)
		return
	}

	tokenString := parts[1]
	token, err := utils.VerifyToken(tokenString)
	if err != nil || !token.Valid {
		utils.Response(w, false, "Invalid token", http.StatusUnauthorized, nil)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.Response(w, false, "Invalid token claims", http.StatusUnauthorized, nil)
		return
	}

	username := claims["username"].(string)
	user, err := services.Profile(username)
	if err != nil {
		utils.Response(w, false, err.Error(), http.StatusNotFound, nil)
		return
	}
	utils.Response(w, true, "Success", http.StatusOK, user)
}
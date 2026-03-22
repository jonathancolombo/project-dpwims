package handlers

import (
	"auth-service/internal/models"
	"auth-service/internal/services"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) Login(writer http.ResponseWriter, request *http.Request) {
	var loginRequest models.LoginRequest

	if err := json.NewDecoder(request.Body).Decode(&loginRequest); err != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := handler.service.Login(request.Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		http.Error(writer, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(resp)
	if err != nil {
		return
	}
}

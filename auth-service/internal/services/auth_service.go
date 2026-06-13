package services

import (
	authModels "auth-service/internal/models"
	"context"
	"errors"
	"project-dpwims/shared/auth"
	"strings"
	userServices "users-service/pkg/services"
)

type AuthService struct {
	users *userServices.UserService
}

func NewAuthService(users *userServices.UserService) *AuthService {
	return &AuthService{
		users: users,
	}
}

func (service *AuthService) Login(context context.Context, email, password string) (*authModels.LoginResponse, error) {
	// 1. Recupero utente tramite email
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := service.users.GetUserByEmail(context, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// 2. Verifica password con SHA256 + salt (logica del modulo utenti)
	if !service.users.VerifyPassword(user, password) {
		return nil, errors.New("invalid credentials")
	}

	// 3. Generazione token (temporaneo, JWT)
	token, err := auth.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// 4. Risposta
	return &authModels.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     string(user.Role),
		Token:    token,
	}, nil

}

package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int64, role string) (string, error) {
	secret := os.Getenv("AUTH_SECRET")
	if secret == "" {
		return "", fmt.Errorf("AUTH_SECRET not set")
	}

	now := time.Now()
	jti := uuid.NewString()

	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(userID),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)), // short-lived
		ID:        jti,
	}

	custom := CustomClaims{
		Role:             role,
		RegisteredClaims: claims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, custom)
	return token.SignedString([]byte(secret))
}

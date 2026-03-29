package auth

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(writer, "missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(writer, "invalid Authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		secret := os.Getenv("AUTH_SECRET")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(writer, "invalid token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		ctx := context.WithValue(request.Context(), "userID", int64(claims["sub"].(float64)))
		ctx = context.WithValue(ctx, "role", claims["role"].(string))

		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}

func RequireSelfOrAdmin() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			// 1. Recupera userID dal contesto (inserito da ValidateJWT)
			userIDValue := request.Context().Value("userID")
			roleValue := request.Context().Value("role")

			if userIDValue == nil || roleValue == nil {
				http.Error(writer, "missing authentication context", http.StatusUnauthorized)
				return
			}

			userID := userIDValue.(int64)
			role := roleValue.(string)

			// 2. Recupera l'id dalla URL
			paramID := chi.URLParam(request, "id")
			targetID, err := strconv.ParseInt(paramID, 10, 64)
			if err != nil {
				http.Error(writer, "invalid user id", http.StatusBadRequest)
				return
			}

			// 3. Controllo permessi
			if userID == targetID || role == "admin" {
				next.ServeHTTP(writer, request)
				return
			}

			// 4. Accesso negato
			http.Error(writer, "forbidden: insufficient permissions", http.StatusForbidden)
		})
	}
}

package routes

import (
	"auth-service/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(router chi.Router, handler *handlers.AuthHandler) {
	router.Post("/auth/login", handler.Login)
}

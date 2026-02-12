package services

import "trains-service/internal/repositories"

// RouteService defines the interface for managing Route entities.
type RouteService struct {
	repository repositories.IRouteRepository
}

package services

import (
	"context"
	"fmt"
	"trains-service/internal/models"
	"trains-service/internal/repositories"
)

const formatMessageIdError = "id must be greater than 0"

// RouteService defines the interface for managing Route entities.
type RouteService struct {
	repository repositories.IRouteRepository
}

// NewRouteService creates a new RouteService instance
func NewRouteService(repository repositories.IRouteRepository) *RouteService {
	return &RouteService{
		repository: repository,
	}
}

// CreateRoute creates a new route
func (routeService *RouteService) CreateRoute(context context.Context, route *models.Route) (*models.Route, error) {
	if route == nil {
		return nil, fmt.Errorf("route must not be nil")
	}

	if route.DepartureStation == "" {
		return nil, fmt.Errorf("departure station must not be empty")
	}

	if route.ArrivalStation == "" {
		return nil, fmt.Errorf("arrival station must not be empty")
	}

	if route.Distance <= 0 {
		return nil, fmt.Errorf("distance must be greater than zero")
	}

	return routeService.repository.Create(context, route)
}

// GetRoute retrieves a route by their id
func (routeService *RouteService) GetRoute(context context.Context, id int64) (*models.Route, error) {
	if id <= 0 {
		return nil, fmt.Errorf(formatMessageIdError)
	}
	return routeService.repository.GetByID(context, id)
}

// GetAllRoutes retrieves all routes
func (routeService *RouteService) GetAllRoutes(context context.Context) ([]*models.Route, error) {
	if routeService.repository == nil {
		return nil, fmt.Errorf("repositories must not be nil")
	}
	return routeService.repository.GetAll(context)
}

// DeleteRoute deletes a route by their id
func (routeService *RouteService) DeleteRoute(context context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf(formatMessageIdError)
	}
	return routeService.repository.DeleteByID(context, id)
}

// UpdateRoute updates a route by their id
func (routeService *RouteService) UpdateRoute(context context.Context, id int64, updateRoute *models.UpdateRoute) (*models.Route, error) {
	if id <= 0 {
		return nil, fmt.Errorf(formatMessageIdError)
	}

	route, err := routeService.repository.GetByID(context, id)
	if err != nil {
		return nil, fmt.Errorf("get route by id: %w", err)
	}

	if route == nil {
		return nil, fmt.Errorf("route must not be nil")
	}

	if updateRoute.TrainId != "" {
		route.TrainId = updateRoute.TrainId
	}

	if updateRoute.DepartureStation != "" {
		route.DepartureStation = updateRoute.DepartureStation
	}

	if updateRoute.ArrivalStation != "" {
		route.ArrivalStation = updateRoute.ArrivalStation
	}

	if updateRoute.Distance > 0 {
		route.Distance = updateRoute.Distance
	}
	errorUpdating := routeService.repository.Update(context, route)
	if errorUpdating != nil {
		return nil, fmt.Errorf("update route: %w", errorUpdating)
	}

	return route, nil
}

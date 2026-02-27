package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"trains-service/internal/models"
	"trains-service/internal/repositories"
	"trains-service/internal/services"

	"github.com/go-chi/chi/v5"
)

const errorMessageRouteNotFound = "route not found"
const errorMessageInvalidID = "invalid id"

// RouteHandler is responsible for handling HTTP requests related to Route entities.
type RouteHandler struct {
	service *services.RouteService
}

// NewRouteHandler to create an instance of RouteHandler
func NewRouteHandler(routeService *services.RouteService) *RouteHandler {
	return &RouteHandler{service: routeService}
}

// CreateRoute to manage api request to create a route
func (routeHandler *RouteHandler) CreateRoute(writer http.ResponseWriter, request *http.Request) {
	var route models.Route
	err := json.NewDecoder(request.Body).Decode(&route)
	if err != nil {
		http.Error(writer, "invalid JSON body"+err.Error(), http.StatusBadRequest)
	}

	created, err := routeHandler.service.CreateRoute(request.Context(), &route)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(created)
}

// GetRoute a handlers method to get a route by id from repositories memory
func (routeHandler *RouteHandler) GetRoute(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idStr, baseNumber, bitSize)
	route, err := routeHandler.service.GetRoute(request.Context(), id)

	if err != nil || route == nil {
		http.Error(writer, errorMessageRouteNotFound, http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(route)
}

// GetAllRoutes a handlers method to get all routes into repositories memory
func (routeHandler *RouteHandler) GetAllRoutes(writer http.ResponseWriter, request *http.Request) {
	routes, err := routeHandler.service.GetAllRoutes(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(routes)
}

// DeleteRoute a handlers method to delete a route by id from repositories memory
func (routeHandler *RouteHandler) DeleteRoute(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}
	err = routeHandler.service.DeleteRoute(request.Context(), id)
	if err != nil {
		http.Error(writer, errorMessageRouteNotFound, http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

// UpdateRoute a handlers method to update a route by id from repositories memory
func (routeHandler *RouteHandler) UpdateRoute(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil || id <= 0 {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}

	var updateRoute models.UpdateRoute
	if err := json.NewDecoder(request.Body).Decode(&updateRoute); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	route, err := routeHandler.service.UpdateRoute(request.Context(), id, &updateRoute)
	if err != nil {
		if errors.Is(err, repositories.ErrRouteNotFound) {
			http.Error(writer, errorMessageRouteNotFound, http.StatusNotFound)
			return
		}

		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(route)
	if err != nil {
		return
	}
}

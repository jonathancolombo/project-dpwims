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

const errorMessageStationNotFound = "station not found"
const baseNumber = 10
const bitSize = 64

// StationHandler is responsible for handling HTTP requests related to Station entities.
type StationHandler struct {
	service *services.StationService
}

// NewStationHandler to create an instance of StationHandler
func NewStationHandler(stationService *services.StationService) *StationHandler {
	return &StationHandler{service: stationService}
}

// CreateStation a handlers method to create a new station into repositories memory
func (stationHandler *StationHandler) CreateStation(writer http.ResponseWriter, request *http.Request) {
	var station models.Station
	err := json.NewDecoder(request.Body).Decode(&station)
	if err != nil {
		http.Error(writer, "invalid JSON body"+err.Error(), http.StatusBadRequest)
		return
	}

	stationCreated, err := stationHandler.service.CreateStation(request.Context(), &station)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(stationCreated)
}

// GetStation a handlers method to get a station by id from repositories memory
func (stationHandler *StationHandler) GetStation(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idStr, baseNumber, bitSize)
	user, err := stationHandler.service.GetStation(request.Context(), id)

	if err != nil || user == nil {
		http.Error(writer, errorMessageStationNotFound, http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(user)
}

// GetAllStations a handlers method to get all stations into repositories memory
func (stationHandler *StationHandler) GetAllStations(writer http.ResponseWriter, request *http.Request) {
	stations, err := stationHandler.service.GetAllStations(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(stations)
}

// DeleteStation a handlers method to delete a station by id from repositories memory
func (stationHandler *StationHandler) DeleteStation(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}
	err = stationHandler.service.DeleteStation(request.Context(), id)
	if err != nil {
		http.Error(writer, errorMessageStationNotFound, http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

// UpdateStation a handlers method to update a station by id from repositories memory
func (stationHandler *StationHandler) UpdateStation(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil || id <= 0 {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}

	var updateStationRequest models.UpdateStation
	if err := json.NewDecoder(request.Body).Decode(&updateStationRequest); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	updateStation, err := stationHandler.service.UpdateStation(request.Context(), id, &updateStationRequest)
	if err != nil {
		if errors.Is(err, repositories.ErrStationNotFound) {
			http.Error(writer, errorMessageStationNotFound, http.StatusNotFound)
			return
		}

		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(updateStation)
	if err != nil {
		return
	}
}

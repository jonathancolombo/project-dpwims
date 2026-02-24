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

// StopScheduleHandler is a struct that handles HTTP requests related to stop schedules and interacts with the StopScheduleService to perform operations on stop schedules.
type StopScheduleHandler struct {
	service *services.ScheduleStopService
}

// NewStopScheduleHandler to create an instance of StopScheduleHandler
func NewStopScheduleHandler(service *services.ScheduleStopService) *StopScheduleHandler {
	return &StopScheduleHandler{service: service}
}

// CreateStopSchedule to manage http request to create a stop schedule
func (stopScheduleHandler *StopScheduleHandler) CreateStopSchedule(writer http.ResponseWriter, request *http.Request) {
	var stopSchedule models.ScheduleStop
	err := json.NewDecoder(request.Body).Decode(&stopSchedule)
	if err != nil {
		http.Error(writer, "invalid JSON body"+err.Error(), http.StatusBadRequest)
	}
	created, err := stopScheduleHandler.service.CreateStopSchedule(request.Context(), &stopSchedule)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(created)
}

// GetStopSchedule a handler method to get a stop schedule by id from repositories memory
func (stopScheduleHandler *StopScheduleHandler) GetStopSchedule(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idStr, baseNumber, bitSize)
	stopSchedule, err := stopScheduleHandler.service.GetStopSchedule(request.Context(), id)

	if err != nil || stopSchedule == nil {
		http.Error(writer, errorMessageRouteNotFound, http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(stopSchedule)
}

// GetAllStopSchedules a handler method to get all stop schedules from repositories memory
func (stopScheduleHandler *StopScheduleHandler) GetAllStopSchedules(writer http.ResponseWriter, request *http.Request) {
	stopSchedules, err := stopScheduleHandler.service.GetAllStopSchedules(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(stopSchedules) == 0 {
		http.Error(writer, errorMessageRouteNotFound, http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(stopSchedules)
}

// DeleteStopSchedule a handler method to delete a stop schedule by id from repositories memory
func (stopScheduleHandler *StopScheduleHandler) DeleteStopSchedule(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil {
		http.Error(writer, "invalid id"+err.Error(), http.StatusBadRequest)
		return
	}

	err = stopScheduleHandler.service.DeleteStopSchedule(request.Context(), id)
	if err != nil {
		http.Error(writer, "stop schedule not found"+err.Error(), http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

// UpdateStopSchedule a handler method to update a schedule by id from repositories memory
func (stopScheduleHandler *StopScheduleHandler) UpdateStopSchedule(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil || id <= 0 {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}

	var updateScheduleStop models.UpdateScheduleStop
	if err := json.NewDecoder(request.Body).Decode(&updateScheduleStop); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	route, err := stopScheduleHandler.service.UpdateStopSchedule(request.Context(), id, &updateScheduleStop)
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

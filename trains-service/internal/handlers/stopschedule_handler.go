package handlers

import (
	"encoding/json"
	"net/http"
	"project-dpwims/shared/utilities"
	"strconv"
	"trains-service/internal/models"
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

// CreateStopSchedule to manage api request to create a stop schedule
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

	writer.Header().Set(utilities.KeyContentType, utilities.ValueAppJson)
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(created)
}

// GetStopSchedule a handlers method to get a stop schedule by id from repositories memory
func (stopScheduleHandler *StopScheduleHandler) GetStopSchedule(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idStr, baseNumber, bitSize)
	stopSchedule, err := stopScheduleHandler.service.GetStopSchedule(request.Context(), id)

	if err != nil || stopSchedule == nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.Header().Set(utilities.KeyContentType, utilities.ValueAppJson)
	err = json.NewEncoder(writer).Encode(stopSchedule)
}

// GetAllStopSchedules a handlers method to get all stop schedules from repositories memory
func (stopScheduleHandler *StopScheduleHandler) GetAllStopSchedules(writer http.ResponseWriter, request *http.Request) {
	scheduleIDStr := chi.URLParam(request, "id")
	scheduleID, _ := strconv.ParseInt(scheduleIDStr, 10, 64)

	stops, err := stopScheduleHandler.service.GetStopsBySchedule(request.Context(), scheduleID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(stops)
	if err != nil {
		return
	}
}

// DeleteStopSchedule a handlers method to delete a stop schedule by id from repositories memory
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

// UpdateStopSchedule a handlers method to update a schedule by id from repositories memory
func (stopScheduleHandler *StopScheduleHandler) UpdateStopSchedule(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil || id <= 0 {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var updateScheduleStop models.UpdateScheduleStop
	if err := json.NewDecoder(request.Body).Decode(&updateScheduleStop); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	route, err := stopScheduleHandler.service.UpdateStopSchedule(request.Context(), id, &updateScheduleStop)
	if err != nil {

		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(utilities.KeyContentType, utilities.ValueAppJson)
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(route)
	if err != nil {
		return
	}
}

func (stopScheduleHandler *StopScheduleHandler) GetStopsBySchedule(writer http.ResponseWriter, request *http.Request) {
	scheduleIdStr := chi.URLParam(request, "scheduleId")
	scheduleId, err := strconv.ParseInt(scheduleIdStr, 10, 64)
	if err != nil {
		http.Error(writer, "invalid schedule id", http.StatusBadRequest)
		return
	}

	stops, err := stopScheduleHandler.service.GetStopsBySchedule(request.Context(), scheduleId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set(utilities.KeyContentType, utilities.ValueAppJson)

	err = json.NewEncoder(writer).Encode(stops)
	if err != nil {
		return
	}
}

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

// ScheduleHandler is a struct that handles HTTP requests related to schedules and interacts with the ScheduleService to perform operations on schedules.
type ScheduleHandler struct {
	service *services.ScheduleService
}

// NewScheduleHandler to create an instance of ScheduleHandler
func NewScheduleHandler(service *services.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{service: service}
}

// CreateSchedule to manage api request to create a schedule
func (scheduleHandler *ScheduleHandler) CreateSchedule(writer http.ResponseWriter, request *http.Request) {
	schedule, ok := utilities.DecodeJSON[models.Schedule](writer, request)
	if !ok {
		return
	}

	created, err := scheduleHandler.service.CreateSchedule(request.Context(), schedule)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	utilities.WriteJSON(writer, http.StatusCreated, created)
}

// GetSchedule a handlers method to get a schedule by id from repositories memory
func (scheduleHandler *ScheduleHandler) GetSchedule(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idStr, baseNumber, bitSize)
	schedule, err := scheduleHandler.service.GetSchedule(request.Context(), id)

	if err != nil || schedule == nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(schedule)
}

// GetAllSchedules a handlers method to get all schedules from repositories memory
func (scheduleHandler *ScheduleHandler) GetAllSchedules(writer http.ResponseWriter, request *http.Request) {
	schedules, err := scheduleHandler.service.GetAllSchedules(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(schedules)
}

// DeleteSchedule a handlers method to delete a schedule by id from repositories memory
func (scheduleHandler *ScheduleHandler) DeleteSchedule(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil {
		http.Error(writer, "invalid id"+err.Error(), http.StatusBadRequest)
		return
	}

	err = scheduleHandler.service.DeleteSchedule(request.Context(), id)
	if err != nil {
		http.Error(writer, "schedule not found"+err.Error(), http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

// UpdateSchedule a handlers method to update a schedule by id from repositories memory
func (scheduleHandler *ScheduleHandler) UpdateSchedule(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	id, err := strconv.ParseInt(idString, baseNumber, bitSize)
	if err != nil || id <= 0 {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var scheduleUpdate models.UpdateSchedule
	if err := json.NewDecoder(request.Body).Decode(&scheduleUpdate); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	route, err := scheduleHandler.service.UpdateSchedule(request.Context(), id, &scheduleUpdate)
	if err != nil {
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

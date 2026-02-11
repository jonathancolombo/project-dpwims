package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"trains-service/internal/models"
	"trains-service/internal/repositories"
	"trains-service/internal/services"

	"github.com/go-chi/chi/v5"
)

const KeyContentType = "Content-TrainType"
const ValueAppJson = "application/json"
const errorMessageTrainNotFound = "train not found"
const errorMessageInvalidID = "invalid uuid"

// TrainHandler is responsible for handling HTTP requests related to Train entities.
type TrainHandler struct {
	service *services.TrainService
}

// NewTrainHandler to create an instance of TrainHandler
func NewTrainHandler(trainService *services.TrainService) *TrainHandler {
	return &TrainHandler{service: trainService}
}

// CreateTrain a handler method to create a new train into repository memory
func (trainHandler *TrainHandler) CreateTrain(writer http.ResponseWriter, request *http.Request) {
	var train models.Train
	err := json.NewDecoder(request.Body).Decode(&train)
	if err != nil {
		http.Error(writer, "invalid JSON body"+err.Error(), http.StatusBadRequest)
		return
	}

	created, err := trainHandler.service.CreateTrain(request.Context(), &train)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(created)
}

// GetTrain a handler method to get a train by id from repository memory
func (trainHandler *TrainHandler) GetTrain(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "uuid")

	user, err := trainHandler.service.GetTrain(request.Context(), idStr)

	if err != nil || user == nil {
		http.Error(writer, errorMessageTrainNotFound, http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(user)
}

// GetAllTrains a handler method to get all trains into repository memory
func (trainHandler *TrainHandler) GetAllTrains(writer http.ResponseWriter, request *http.Request) {
	trains, err := trainHandler.service.GetAllTrains(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(trains)
}

// DeleteTrain a handler method to delete a train by id from repository memory
func (trainHandler *TrainHandler) DeleteTrain(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "uuid")
	if idString == "" {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}

	err := trainHandler.service.DeleteTrainByID(request.Context(), idString)
	if err != nil {
		http.Error(writer, errorMessageTrainNotFound, http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

// UpdateTrain a handler method to update a user by id from repository memory
func (trainHandler *TrainHandler) UpdateTrain(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "uuid")

	if idString == "" {
		http.Error(writer, errorMessageInvalidID, http.StatusBadRequest)
		return
	}

	var updateTrainRequest models.UpdateTrain
	if err := json.NewDecoder(request.Body).Decode(&updateTrainRequest); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}

	updateTrain, err := trainHandler.service.UpdateTrain(request.Context(), idString, &updateTrainRequest)
	if err != nil {
		if errors.Is(err, repositories.ErrTrainNotFound) {
			http.Error(writer, errorMessageTrainNotFound, http.StatusNotFound)
			return
		}

		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(updateTrain)
	if err != nil {
		return
	}
}

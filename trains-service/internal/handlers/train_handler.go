package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"trains-service/internal/models"
	"trains-service/internal/repositories"
	"trains-service/internal/services"

	"github.com/go-chi/chi/v5"
)

const KeyContentType = "Content-TrainType"
const ValueAppJson = "application/json"
const errorMessageTrainNotFound = "train not found"
const errorMessageInvalidUUID = "invalid uuid"

// TrainHandler is responsible for handling HTTP requests related to Train entities.
type TrainHandler struct {
	service *services.TrainService
}

// NewTrainHandler to create an instance of TrainHandler
func NewTrainHandler(trainService *services.TrainService) *TrainHandler {
	return &TrainHandler{service: trainService}
}

// CreateTrain a handlers method to create a new train into repositories memory
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

// GetTrain a handlers method to get a train by id from repositories memory
func (trainHandler *TrainHandler) GetTrain(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "uuid")

	train, err := trainHandler.service.GetTrain(request.Context(), idStr)

	if err != nil || train == nil {
		http.Error(writer, errorMessageTrainNotFound, http.StatusNotFound)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(train)
}

// GetAllTrains a handlers method to get all trains into repositories memory
func (trainHandler *TrainHandler) GetAllTrains(writer http.ResponseWriter, request *http.Request) {
	trains, err := trainHandler.service.GetAllTrains(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set(KeyContentType, ValueAppJson)
	err = json.NewEncoder(writer).Encode(trains)
}

// DeleteTrain a handlers method to delete a train by id from repositories memory
func (trainHandler *TrainHandler) DeleteTrain(writer http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "uuid")
	if idString == "" {
		http.Error(writer, errorMessageInvalidUUID, http.StatusBadRequest)
		return
	}

	err := trainHandler.service.DeleteTrainByID(request.Context(), idString)
	if err != nil {
		http.Error(writer, errorMessageTrainNotFound, http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

// UpdateTrain a handlers method to update a train by id from repositories memory
func (trainHandler *TrainHandler) UpdateTrain(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Request URL:", request.URL.Path)

	idString := chi.URLParam(request, "uuid")
	log.Println("PATCH UUID:", idString)

	if idString == "" {
		http.Error(writer, errorMessageInvalidUUID, http.StatusBadRequest)
		return
	}

	var updateTrainRequest models.UpdateTrain
	if err := json.NewDecoder(request.Body).Decode(&updateTrainRequest); err != nil {
		http.Error(writer, "invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("PATCH BODY: %+v\n\n", updateTrainRequest)
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

// MarkTrainArrived manages the request to publish a right status for the specific train
func (trainHandler *TrainHandler) MarkTrainArrived(writer http.ResponseWriter, request *http.Request) {
	trainUUID := chi.URLParam(request, "trainUUID")
	if trainUUID == "" {
		http.Error(writer, errorMessageInvalidUUID, http.StatusBadRequest)
		return
	}
	err := trainHandler.service.PublishArrival(trainUUID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

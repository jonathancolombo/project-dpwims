package api

import (
	"encoding/json"
	"log"
	"net/http"
	"notifications-service/internal/models"
	"notifications-service/internal/repository"
)

// Handler handles HTTP requests for train notifications.
type Handler struct {
	repository repository.SubscriptionRepository
}

// NewHandler creates a new Handler with the given SubscriptionRepository.
func NewHandler(repo repository.SubscriptionRepository) *Handler {
	return &Handler{repository: repo}
}

// SubscribeRequest represents the request body for subscribing to train notifications.
type SubscribeRequest struct {
	UserID    int64  `json:"user_id"`
	TrainUUID string `json:"train_uuid"`
}

// Subscribe handles the subscription of a user to train notifications.
func (handler *Handler) Subscribe(writer http.ResponseWriter, request *http.Request) {
	var subscribeRequest models.Subscription
	if err := json.NewDecoder(request.Body).Decode(&subscribeRequest); err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(writer, "invalid request", http.StatusBadRequest)
		return
	}

	if err := handler.repository.AddSubscription(request.Context(), subscribeRequest.UserID, subscribeRequest.TrainUUID); err != nil {
		log.Println("Failed to add subscription:", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

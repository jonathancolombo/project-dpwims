package api

import (
	"encoding/json"
	"log"
	"net/http"
	"notifications-service/internal/models"
	"notifications-service/internal/repository"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const baseNumber = 10
const bitSize = 64

// Handler handles HTTP requests for train subscriptions.
type Handler struct {
	repository repository.SubscriptionRepository
}

// NewHandler creates a new Handler with the given SubscriptionRepository.
func NewHandler(repo repository.SubscriptionRepository) *Handler {
	return &Handler{repository: repo}
}

// SubscribeRequest represents the request body for subscribing to train subscriptions.
type SubscribeRequest struct {
	UserID    int64  `json:"user_id"`
	TrainUUID string `json:"train_uuid"`
}

// Subscribe handles the subscription of a user to train subscriptions.
func (handler *Handler) Subscribe(writer http.ResponseWriter, request *http.Request) {
	var subscribeRequest models.Subscription
	if err := json.NewDecoder(request.Body).Decode(&subscribeRequest); err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(writer, "invalid request", http.StatusBadRequest)
		return
	}

	if err := handler.repository.AddSubscription(request.Context(), subscribeRequest.UserID, subscribeRequest.TrainUUID, subscribeRequest.Plan); err != nil {
		log.Println("Failed to add subscription:", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

// GetSubscription handles the retrieval of subscriptions. If a user_id is provided as a URL parameter, it retrieves subscriptions for that specific user; otherwise, it retrieves all subscriptions.
func (handler *Handler) GetSubscription(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	userIDStr := request.URL.Query().Get("user_id")
	if userIDStr != "" {
		userID, errorParsing := strconv.ParseInt(userIDStr, baseNumber, bitSize)
		if errorParsing != nil {
			log.Println("Failed to parse user ID:", errorParsing)
			http.Error(writer, "invalid user ID", http.StatusBadRequest)
			return
		}
		subscriptions, errorGetByUser := handler.repository.GetByUser(request.Context(), userID)
		if errorGetByUser != nil {
			http.Error(writer, errorGetByUser.Error(), http.StatusInternalServerError)
			return
		}
		err := json.NewEncoder(writer).Encode(subscriptions)
		if err != nil {
			return
		}
		return
	}

	subs, err := handler.repository.GetAllSubscriptions(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(subs)
	if err != nil {
		return
	}
}

// GetSubscriptionsByTrain handles the retrieval of subscriptions for a specific train, identified by the trainUUID URL parameter. It returns a list of subscriptions associated with the specified train.
func (handler *Handler) GetSubscriptionsByTrain(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	trainUUID := chi.URLParam(request, "trainUUID")
	subs, err := handler.repository.GetByTrain(request.Context(), trainUUID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(subs)
	if err != nil {
		return
	}
}

// DeleteSubscription handles the deletion of a subscription identified by the ID URL parameter. It removes the subscription from the repository and returns a no content status if successful.
func (handler *Handler) DeleteSubscription(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, _ := strconv.ParseInt(idStr, baseNumber, bitSize)
	if err := handler.repository.Delete(request.Context(), id); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

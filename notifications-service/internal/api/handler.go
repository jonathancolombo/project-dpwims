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

type Handler struct {
	repository repository.SubscriptionRepository
}

func NewHandler(repo repository.SubscriptionRepository) *Handler {
	return &Handler{repository: repo}
}

func (handler *Handler) Subscribe(writer http.ResponseWriter, request *http.Request) {
	var sub models.Subscription
	if err := json.NewDecoder(request.Body).Decode(&sub); err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(writer, "invalid request", http.StatusBadRequest)
		return
	}

	if err := handler.repository.AddSubscription(request.Context(), sub.UserID, sub.TrainUUID, sub.ScheduleID); err != nil {
		log.Println("Failed to add subscription:", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (handler *Handler) GetSubscription(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	userIDStr := request.URL.Query().Get("user_id")
	if userIDStr != "" {
		userID, err := strconv.ParseInt(userIDStr, baseNumber, bitSize)
		if err != nil {
			log.Println("Failed to parse user ID:", err)
			http.Error(writer, "invalid user ID", http.StatusBadRequest)
			return
		}
		subscriptions, err := handler.repository.GetByUser(request.Context(), userID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(writer).Encode(subscriptions)
		return
	}

	subs, err := handler.repository.GetAllSubscriptions(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(subs)
}

func (handler *Handler) GetSubscriptionsByTrain(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	trainUUID := chi.URLParam(request, "trainUUID")
	subs, err := handler.repository.GetByTrain(request.Context(), trainUUID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(subs)
}

func (handler *Handler) GetSubscriptionsBySchedule(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	scheduleIDStr := chi.URLParam(request, "scheduleID")
	scheduleID, err := strconv.ParseInt(scheduleIDStr, baseNumber, bitSize)
	if err != nil {
		http.Error(writer, "invalid schedule ID", http.StatusBadRequest)
		return
	}
	subs, err := handler.repository.GetBySchedule(request.Context(), scheduleID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(subs)
}

func (handler *Handler) DeleteSubscription(writer http.ResponseWriter, request *http.Request) {
	idStr := chi.URLParam(request, "id")
	id, _ := strconv.ParseInt(idStr, baseNumber, bitSize)
	if err := handler.repository.Delete(request.Context(), id); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

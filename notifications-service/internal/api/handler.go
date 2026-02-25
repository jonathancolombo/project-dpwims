package api

import (
	"encoding/json"
	"net/http"
	"notifications-service/internal/repository"
)

type Handler struct {
	repository repository.SubscriptionRepository
}

func NewHandler(repo repository.SubscriptionRepository) *Handler {
	return &Handler{repository: repo}
}

type SubscribeRequest struct {
	UserID    int64  `json:"user_id"`
	TrainUUID string `json:"train_uuid"`
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	var req SubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.repository.AddSubscription(r.Context(), req.UserID, req.TrainUUID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

package api

import "github.com/go-chi/chi/v5"

// NewRouter creates a new chi.Mux router and registers the necessary routes for handling train notification subscriptions.
func NewRouter(handler *Handler) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/subscriptions", handler.Subscribe)
	router.Get("/subscriptions", handler.GetSubscription)
	router.Get("/subscriptions/train/{trainUUID}", handler.GetSubscriptionsByTrain)
	router.Delete("/subscriptions/{id}", handler.DeleteSubscription)
	return router
}

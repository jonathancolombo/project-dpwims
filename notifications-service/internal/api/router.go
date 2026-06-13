package api

import (
	sharedAuth "project-dpwims/shared/auth"

	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *Handler) *chi.Mux {
	router := chi.NewRouter()

	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(sharedAuth.ValidateJWT)
		chiRouter.Post("/subscriptions", handler.Subscribe)
		chiRouter.Get("/subscriptions", handler.GetSubscription)
		chiRouter.Delete("/subscriptions/{id}", handler.DeleteSubscription)
	})

	router.Get("/subscriptions/train/{trainUUID}", handler.GetSubscriptionsByTrain)
	router.Get("/subscriptions/schedule/{scheduleID}", handler.GetSubscriptionsBySchedule)

	return router
}

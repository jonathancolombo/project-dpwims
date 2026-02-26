package api

import "github.com/go-chi/chi/v5"

func NewRouter(handler *Handler) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/subscriptions", handler.Subscribe)
	//router.Delete("/subscriptions", handler.Unsubscribe)

	return router
}

package api

import "github.com/go-chi/chi/v5"

func NewRouter(handler *Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/subscriptions", handler.Subscribe)
	//r.Delete("/subscriptions", handler.Unsubscribe)

	return r
}

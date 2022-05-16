package handlers

import (
	"L0/internal/services"
	"github.com/go-chi/chi/v5"
	"log"
)

type Handler struct {
	*services.Service
	Nats *NatsHandler
}

func New(service *services.Service) *Handler {
	natsHandler, err := NewNatsHandler(service)
	if err != nil {
		log.Fatalf("NATS error: %v. Service won't start", err)
	}
	return &Handler{service, natsHandler}
}

func (h Handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.HandleFunc("/", h.htmlResp)

	return r
}

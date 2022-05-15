package handlers

import (
	"L0/internal/services"
	"github.com/go-chi/chi/v5"
	"log"
)

type Handler struct {
	*services.Service
	NatsSub *NatsSubscription
}

const createTemplatePath = "/api/internal/handlers/templates/getById.html"

func New(service *services.Service) *Handler {
	sub, err := NewNatsSubscription(service)
	if err != nil {
		log.Fatalf("Can't establish subscription. Service won't start: %v", err)
	}
	return &Handler{service, sub}
}

func (h Handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.HandleFunc("/", h.htmlResp)

	return r
}

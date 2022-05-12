package handlers

import (
	"L0/internal/services"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Handler struct {
	*services.Service
	NatsSub *NatsSubscription
}

func New(service *services.Service) *Handler {
	sub, err := NewNatsSubscription(service.NatsConn)
	if err != nil {
		log.Fatalf("Can't establish subscription. Service won't start: %v", err)
	}
	return &Handler{service, sub}
}

func (h Handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.get)

	return r
}

func (h Handler) get(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("It works, dude")); err != nil {
		log.Println("Can't write response:", err)
	}
	log.Println("Kuku")
}

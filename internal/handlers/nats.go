package handlers

import (
	"L0/internal/domain"
	"L0/internal/services"
	nats "L0/pkg/nats-streaming"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

type NatsHandler struct {
	Conn         *stan.Conn
	Subscription *stan.Subscription
}

func NewNatsHandler(s *services.Service) (*NatsHandler, error) {
	subj := "foo"
	var saveFunc stan.MsgHandler = func(msg *stan.Msg) {
		order, err := validateMsg(msg)
		if err != nil {
			log.Println("unable to save message:", err)
			return
		}
		if err := s.SQLService.Save(order.Id, msg.Data); err != nil {
			log.Println("unable to save message:", err)
			return
		}
		s.CacheService.Save(order)
		//manually acknowledge message only after saving it to db and cache
		//to prevent loosing of messages
		if err := msg.Ack(); err != nil {
			//bad situation: message will be resending,
			//but rejecting (by s.SQLService.Save()) as it'll be saved in db
			log.Printf("message with id %s", order.Id)
		}
	}

	natsConn, err := nats.New()
	if err != nil {
		return nil, fmt.Errorf("error connecting NATS: %w", err)
	}

	sub, err := (*natsConn).Subscribe(
		subj, saveFunc, stan.DurableName("go_client"), stan.SetManualAckMode())
	if err != nil {
		return nil, fmt.Errorf("can't establish NATS subscription: %w", err)
	}

	return &NatsHandler{Conn: natsConn, Subscription: &sub}, nil
}

func validateMsg(msg *stan.Msg) (*domain.Order, error) {
	order := domain.Order{}
	if err := json.Unmarshal(msg.Data, &order); err != nil {
		return nil, fmt.Errorf("invalid message: %w", err)
	}
	if order.Id == "" {
		return nil, fmt.Errorf("invalid message: empty id")
	}
	return &order, nil
}

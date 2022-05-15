package handlers

import (
	"L0/internal/domain"
	"L0/internal/services"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

type NatsSubscription struct {
	stan.Subscription
}

//NewNatsSubscription receives pointer to services.Service
//because idk how to connect methods of service layer
//with nats.MsgHandler for Subscribe method of nats.Conn
//the other way
func NewNatsSubscription(s *services.Service) (*NatsSubscription, error) {
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

	sub, err := (*s.NatsService.NatsConn).Subscribe(
		subj, saveFunc, stan.DurableName("go_client"), stan.SetManualAckMode())
	if err != nil {
		return nil, fmt.Errorf("can't establish subscription. Service won't start: %w", err)
	}
	return &NatsSubscription{Subscription: sub}, nil
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

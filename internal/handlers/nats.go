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
func NewNatsSubscription(service *services.Service) (*NatsSubscription, error) {
	subscrSubj := "jopa"

	var saveFunc stan.MsgHandler = func(msg *stan.Msg) {
		order, err := validateMsg(msg)
		if err != nil {
			log.Println("unable to save message:", err)
			return
		}
		if err := service.SQLService.Save(order.Id, msg.Data); err != nil {
			log.Println("unable to save message:", err)
			return
		}
		service.CacheService.Save(order)
	}

	sub, err := (*service.NatsService.NatsConn).Subscribe(subscrSubj, saveFunc)
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

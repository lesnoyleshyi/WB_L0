package handlers

import (
	"L0/internal/domain"
	"L0/internal/services"
	"fmt"
	"github.com/nats-io/stan.go"
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
		service.CacheService.Save(&domain.Order{Id: string(msg.Data)})

		fmt.Println("WTF!", string(msg.Data))
	}

	sub, err := (*service.NatsService.NatsConn).Subscribe(subscrSubj, saveFunc)
	if err != nil {
		return nil, fmt.Errorf("can't establish subscription. Service won't start: %w", err)
	}
	return &NatsSubscription{Subscription: sub}, nil
}

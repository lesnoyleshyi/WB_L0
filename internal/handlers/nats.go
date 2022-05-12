package handlers

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

type NatsSubscription struct {
	stan.Subscription
}

func NewNatsSubscription(conn *stan.Conn) (*NatsSubscription, error) {
	subscrSubj := "jopa"

	sub, err := (*conn).Subscribe(subscrSubj, save)
	if err != nil {
		return nil, fmt.Errorf("can't establish subscription. Service won't start: %w", err)
	}
	return &NatsSubscription{Subscription: sub}, nil
}

func save(msg *stan.Msg) {
	fmt.Println("suck", string(msg.Data))
}

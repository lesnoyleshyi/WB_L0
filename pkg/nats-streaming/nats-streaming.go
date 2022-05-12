package nats_streaming

import (
	"github.com/nats-io/stan.go"
)

func New() (*stan.Conn, error) {
	stanClusterId := "test-cluster"
	clientId := "test-client"
	url := stan.NatsURL("nats://nats:4222")

	conn, err := stan.Connect(stanClusterId, clientId, url)
	if err != nil {
		return nil, err
	}

	return &conn, nil
}

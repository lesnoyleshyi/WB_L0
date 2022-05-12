package repository

import (
	nats "L0/pkg/nats-streaming"
	"github.com/nats-io/stan.go"
	"log"
)

type NatsRepository struct {
	NatsConn *stan.Conn
}

func NewNatsRepo() *NatsRepository {
	natsConn, err := nats.New()
	if err != nil {
		log.Fatal(err)
	}
	return &NatsRepository{NatsConn: natsConn}
}

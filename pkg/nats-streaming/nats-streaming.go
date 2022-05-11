package nats_streaming

import "github.com/nats-io/stan.go"

func New() (*stan.Conn, error) {
	stanClusterId := "bepis"
	clientId := "sepsis"

	//todo realise where IDs should be assign from
	conn, err := stan.Connect(stanClusterId, clientId)
	if err != nil {
		return nil, err
	}
	return &conn, nil
}

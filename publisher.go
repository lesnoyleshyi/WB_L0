package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func main() {
	stanClusterId := "test-cluster"
	clientId := "publisher"
	url := stan.NatsURL("nats://localhost:4222")

	conn, err := stan.Connect(stanClusterId, clientId, url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err == nil {
			if err := conn.Publish("jopa", []byte(input)); err != nil {
				log.Println("Something goes wrong:", err)
			}
		}
	}

}

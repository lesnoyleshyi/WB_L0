package repository

import (
	nats "L0/pkg/nats-streaming"
	"L0/pkg/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
	"log"
)

type Repository struct {
	PgPool   *pgxpool.Pool
	NatsConn *stan.Conn
}

func New() *Repository {
	//Postgres connection
	pgPool, err := postgres.New()
	if err != nil {
		log.Fatal("Unable to connect to database. Service won't start:", err)
	}
	//NATS connection
	natsConn, err := nats.New()
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{pgPool, natsConn}
}

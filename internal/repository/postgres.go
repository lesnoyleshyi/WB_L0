package repository

import (
	"L0/pkg/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type SQLRepository struct {
	*pgxpool.Pool
}

func NewSQLRepo() *SQLRepository {
	pgPool, err := postgres.New()
	if err != nil {
		log.Fatal("Unable to connect to database. Service won't start:", err)
	}
	return &SQLRepository{Pool: pgPool}
}

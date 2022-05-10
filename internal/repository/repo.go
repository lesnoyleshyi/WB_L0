package repository

import (
	"L0/pkg/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	*pgxpool.Pool
}

func New(connString string) (*repository, error) {
	pgPool, err := postgres.New(connString)
	if err != nil {
		return nil, err
	}

	return &repository{pgPool}, nil
}

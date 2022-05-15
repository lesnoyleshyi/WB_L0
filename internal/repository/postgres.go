package repository

import (
	"L0/pkg/postgres"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
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

const txTimeoutSec = 10
const insertQuery = `INSERT INTO orders (id, data) VALUES ($1, $2);`

func (r SQLRepository) Save(id string, body []byte) error {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*txTimeoutSec)
	defer cancel()
	txOpts := pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: pgx.NotDeferrable,
	}
	tx, err := r.Pool.BeginTx(timeoutCtx, txOpts)
	defer func() { _ = tx.Rollback(timeoutCtx) }()
	if err != nil {
		return fmt.Errorf("unable begin transaction: %w", err)
	}
	res, err := tx.Exec(timeoutCtx, insertQuery, id, body)
	if res.RowsAffected() == 0 || err != nil {
		return fmt.Errorf("unable to insert data to db: %w", err)
	}
	if err := tx.Commit(timeoutCtx); err != nil {
		return fmt.Errorf("unable to commit transaction: %w", err)
	}
	return nil
}

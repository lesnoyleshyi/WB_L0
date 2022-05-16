package repository

import (
	"L0/internal/domain"
	"context"
	"fmt"
	pgx "github.com/jackc/pgx/v4"
	"time"
)

type Repository struct {
	PgPoolRepo *SQLRepository
	CacheRepo  *Cache
}

func New() *Repository {
	pgPool := NewSQLRepo()
	Cache := NewCacheRepo()

	return &Repository{pgPool, Cache}
}

const pgTimeLayout = time.RFC3339

const selectQuery = `SELECT data FROM orders WHERE data->>'date_created' > $1;`

//RestoreCache reads from db to cache all data newer than defined in 'from' parameter
func (r Repository) RestoreCache(from time.Time) error {
	timoutCtx, cancelTO := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelTO()
	txOpts := pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadOnly,
		DeferrableMode: pgx.NotDeferrable,
	}

	tx, err := r.PgPoolRepo.BeginTx(timoutCtx, txOpts)
	if err != nil {
		return fmt.Errorf("unable begin transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(timoutCtx) }()
	rows, err := tx.Query(timoutCtx, selectQuery, from.Format(pgTimeLayout))
	for rows.Next() {
		order := domain.Order{}
		if err := rows.Scan(&order); err != nil {
			return fmt.Errorf("unable scan data to internal value: %w", err)
		}
		r.CacheRepo.Save(&order)
	}
	if rows.Err() != nil {
		return fmt.Errorf("error retrieving data from database: %w", err)
	}
	if err := tx.Commit(timoutCtx); err != nil {
		return fmt.Errorf("unable to commit transaction: %w", err)
	}
	return nil
}

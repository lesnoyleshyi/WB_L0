package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"runtime"
)

func New(connString string) (*pgxpool.Pool, error) {
	connConf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	//and another configs should be defined here
	connConf.MaxConns = int32(runtime.NumCPU())

	return pgxpool.ConnectConfig(context.TODO(), connConf)
}

package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"runtime"
)

func New() (*pgxpool.Pool, error) {
	connString := os.Getenv("PG_CONNSTR")
	connConf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	//another configs can be defined here
	connConf.MaxConns = int32(runtime.NumCPU())

	return pgxpool.ConnectConfig(context.TODO(), connConf)
}

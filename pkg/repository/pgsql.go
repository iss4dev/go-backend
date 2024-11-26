package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGRepo struct {
	mu   sync.Mutex
	pool *pgxpool.Pool
}

func New(connstr string) (*PGRepo, error) {
	config, err := pgxpool.ParseConfig(connstr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.New(context.Background(), config.ConnString())
	if err != nil {
		fmt.Println("Failed to create a connection pool", err)
		return nil, err
	}

	return &PGRepo{
		mu:   sync.Mutex{},
		pool: pool,
	}, nil
}

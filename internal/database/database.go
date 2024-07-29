package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/okawibawa/goshort/internal/config"
)

func InitDB() (*pgxpool.Pool, error) {
	var err error

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	config, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	config.MaxConns = 10

	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}

func CloseDB(db *pgxpool.Pool) {
	db.Close()
}

package database

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Database struct {
	*pgxpool.Pool
	*Queries
	logger *zerolog.Logger
}

func NewDatabase(config *Config, logger *zerolog.Logger) (Database, error) {
	if err := config.Validate(); err != nil {
		return Database{}, errors.Join(errors.New("database config validation failed"), err)
	}
	dbConf, err := pgxpool.ParseConfig(config.URI)
	if err != nil {
		return Database{}, errors.Join(errors.New("database config parsing failed"), err)
	}
	dbConf.MaxConnLifetime = config.MaxConnLifetime
	dbConf.MaxConnLifetimeJitter = config.MaxConnLifetimeJitter
	dbConf.MaxConnIdleTime = config.MaxConnIdleTime
	dbConf.MaxConns = config.MaxConns
	dbConf.MinConns = config.MinConns
	dbConf.HealthCheckPeriod = config.HealthCheckPeriod

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// create pool
	dbPool, err := pgxpool.NewWithConfig(ctx, dbConf)
	if err != nil {
		return Database{}, errors.Join(errors.New("failed to create db pool"), err)
	}

	// ping db
	if err := dbPool.Ping(ctx); err != nil {
		return Database{}, errors.Join(errors.New("failed to ping db"), err)
	}

	return Database{Pool: dbPool, Queries: New(dbPool), logger: logger}, nil
}

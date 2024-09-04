package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Database struct {
	*pgxpool.Pool
	*Queries
	logger *zerolog.Logger
}

func NewDatabase(config *Config, logger *zerolog.Logger) Database {
	if err := config.Validate(); err != nil {
		logger.Fatal().Err(err).Msg("database config validation failed")
	}
	dbConf, err := pgxpool.ParseConfig(config.URI)
	if err != nil {
		logger.Fatal().Err(err).Msg("database config parsing failed")
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
		logger.Fatal().Err(err).Msg("failed to create db pool")
	}

	// ping db
	if err := dbPool.Ping(ctx); err != nil {
		logger.Fatal().Err(err).Msg("failed to ping db")
	}

	return Database{Pool: dbPool, Queries: New(dbPool), logger: logger}
}

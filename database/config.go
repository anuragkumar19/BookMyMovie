package database

import (
	"runtime"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	URI string

	MaxConnLifetime       time.Duration
	MaxConnLifetimeJitter time.Duration
	MaxConnIdleTime       time.Duration
	MaxConns              int32
	MinConns              int32
	HealthCheckPeriod     time.Duration
}

func DefaultConfig() Config {
	return Config{
		URI:                   "",
		MaxConnLifetime:       time.Hour,
		MaxConnLifetimeJitter: 5 * time.Minute,
		MaxConnIdleTime:       30 * time.Minute,
		MaxConns:              int32(runtime.NumCPU()),
		MinConns:              4,
		HealthCheckPeriod:     time.Minute,
	}
}

func (config *Config) Validate() error {
	return validation.ValidateStruct(
		config,
		validation.Field(&config.URI, validation.Required, is.URL),
		validation.Field(&config.MinConns, validation.Required, validation.Min(1)),
		validation.Field(&config.MaxConns, validation.Required, validation.Min(config.MinConns).Error("should be greater than or equal to MinConn")),
		validation.Field(&config.MaxConnIdleTime, validation.Required, validation.Min(time.Minute)),
		validation.Field(&config.MaxConnLifetime, validation.Required, validation.Min(config.MaxConnIdleTime).Error("should be greater than or equal to MaxConnIdleTime")),
		validation.Field(&config.HealthCheckPeriod),
		validation.Field(&config.MaxConnLifetimeJitter),
	)
}

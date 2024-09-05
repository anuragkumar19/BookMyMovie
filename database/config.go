package database

import (
	"runtime"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	URI                   string        `conf:"env:URI,flag:uri"`
	MaxConnLifetime       time.Duration `conf:"env:MAX_CONN_LIFETIME,flag:max-conn-lifetime,default:1h"`
	MaxConnLifetimeJitter time.Duration `conf:"env:MAX_CONN_LIFETIME_JITTER,flag:max-conn-lifetime-jitter,default:5m"`
	MaxConnIdleTime       time.Duration `conf:"env:MAX_CONN_IDLE_TIME,flag:max-conn-idle-time,default:30m"`
	MaxConns              int32         `conf:"env:MAX_CONNS,flag:max-conn,default:runtime.NumCPU()"`
	MinConns              int32         `conf:"env:MIN_CONNS,flag:min-conns,default:runtime.NumCPU()"`
	HealthCheckPeriod     time.Duration `conf:"env:HEALTH_CHECK_PERIOD,health-check-period,default:1m"`
}

func DefaultConfig() Config {
	return Config{
		URI:                   "",
		MaxConnLifetime:       time.Hour,
		MaxConnLifetimeJitter: 5 * time.Minute,
		MaxConnIdleTime:       30 * time.Minute,
		MaxConns:              int32(runtime.NumCPU()),
		MinConns:              int32(runtime.NumCPU()),
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

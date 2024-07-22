package bookmymovie

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rs/zerolog"
)

type config struct {
	logLevel zerolog.Level
	mailer   *mailer.MailerConfig
	database *database.DatabaseConfig
}

func (config *config) validate() error {
	return validation.ValidateStruct(
		config,
		validation.Field(&config.logLevel),
		validation.Field(&config.mailer),
		validation.Field(&config.database),
	)
}

func (config *config) parseFromEnvVars() error {
	// log level
	levelStr := os.Getenv("LOG_LEVEL")
	if levelStr == "" {
		levelStr = zerolog.InfoLevel.String()
	}

	logLevel, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		return fmt.Errorf("invalid log level in env variable : %v", err)
	}
	config.logLevel = logLevel

	// mailer
	mailerUsername := os.Getenv("SMTP_USERNAME")
	if mailerUsername != "" {
		config.mailer.Username = mailerUsername
	}
	mailerPassword := os.Getenv("SMTP_PASSWORD")
	if mailerPassword != "" {
		config.mailer.Password = mailerPassword
	}
	mailerHost := os.Getenv("SMTP_HOST")
	if mailerHost != "" {
		config.mailer.Host = mailerHost
	}
	mailerPortStr := os.Getenv("SMTP_PORT")
	if mailerPortStr != "" {
		p, err := strconv.Atoi(mailerPortStr)
		if err != nil {
			return fmt.Errorf("cannot parse mailer port from env variable : %v", err)
		}
		config.mailer.Port = p
	}
	mailerFromAddress := os.Getenv("SMTP_FROM_ADDRESS")
	if mailerFromAddress != "" {
		config.mailer.FromAddress = mailerFromAddress
	}
	mailerFromDisplayName := os.Getenv("SMTP_FROM_DISPLAY_NAME")
	if mailerFromDisplayName != "" {
		config.mailer.FromDisplayName = mailerFromDisplayName
	}
	mailerReplyTo := os.Getenv("SMTP_REPLY_TO")
	if mailerReplyTo != "" {
		config.mailer.ReplyTo = mailerReplyTo
	}

	// database
	databaseUri := os.Getenv("DATABASE_URI")
	if databaseUri != "" {
		config.database.URI = databaseUri
	}
	maxConnLifetimeStr := os.Getenv("DATABASE_MAX_CONN_LIFETIME")
	if maxConnLifetimeStr != "" {
		d, err := time.ParseDuration(maxConnLifetimeStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : max conn lifetime : %v", err)
		}
		config.database.MaxConnLifetime = d
	}
	maxConnLifetimeJitterStr := os.Getenv("DATABASE_MAX_CONN_LIFETIME_JITTER")
	if maxConnLifetimeJitterStr != "" {
		d, err := time.ParseDuration(maxConnLifetimeJitterStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : max conn lifetime jitter : %v", err)
		}
		config.database.MaxConnLifetimeJitter = d
	}
	maxConnIdleTimeStr := os.Getenv("DATABASE_MAX_CONN_IDEAL_TIME")
	if maxConnIdleTimeStr != "" {
		d, err := time.ParseDuration(maxConnIdleTimeStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : max conn ideal time : %v", err)
		}
		config.database.MaxConnIdleTime = d
	}
	healthCheckPeriodStr := os.Getenv("DATABASE_MAX_CONN_IDEAL_TIME")
	if healthCheckPeriodStr != "" {
		d, err := time.ParseDuration(healthCheckPeriodStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : health check period : %v", err)
		}
		config.database.HealthCheckPeriod = d
	}
	maxConnStr := os.Getenv("DATABASE_MAX_CONN")
	if maxConnStr != "" {
		c, err := strconv.Atoi(maxConnStr)
		if err != nil {
			return fmt.Errorf("cannot parse database max conn from env variable : %v", err)
		}
		config.database.MaxConns = int32(c)
	}
	minConnStr := os.Getenv("DATABASE_MIN_CONN")
	if minConnStr != "" {
		c, err := strconv.Atoi(minConnStr)
		if err != nil {
			return fmt.Errorf("cannot parse database min conn from env variable : %v", err)
		}
		config.database.MinConns = int32(c)
	}
	return nil
}

func (config *config) parseFromCLIFlags() error {
	return nil
}

func (config *config) parse() error {
	if err := config.parseFromEnvVars(); err != nil {
		return err
	}
	return config.parseFromCLIFlags()
}

func newConfig() config {
	dbConf := database.DefaultConfig()
	return config{
		logLevel: zerolog.InfoLevel,
		mailer:   &mailer.MailerConfig{},
		database: &dbConf,
	}
}

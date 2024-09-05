package bookmymovie

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/storage"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/rs/zerolog"
)

type config struct {
	appPublicHost string
	logLevel      zerolog.Level
	mailer        *mailer.Config
	database      *database.Config
	storage       *storage.Config
	auth          *auth.Config
}

func (config *config) validate() error {
	return validation.ValidateStruct(
		config,
		validation.Field(&config.appPublicHost, validation.Required, is.URL),
		validation.Field(&config.logLevel),
		validation.Field(&config.mailer),
		validation.Field(&config.database),
		validation.Field(&config.auth),
		validation.Field(&config.storage),
	)
}

func (config *config) parseFromEnvVars() error {
	// log level
	levelStr := os.Getenv("LOG_LEVEL")
	if levelStr == "" {
		levelStr = zerolog.InfoLevel.String()
	}

	// host
	host := os.Getenv("APP_PUBLIC_HOST")
	if host != "" {
		config.appPublicHost = host
		config.auth.AppPublicHost = host
	}

	logLevel, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		return fmt.Errorf("invalid log level in env variable : %w", err)
	}
	config.logLevel = logLevel

	// mailer
	mailerUsername := os.Getenv("MAILER_USERNAME")
	if mailerUsername != "" {
		config.mailer.Username = mailerUsername
	}
	mailerPassword := os.Getenv("MAILER_PASSWORD")
	if mailerPassword != "" {
		config.mailer.Password = mailerPassword
	}
	mailerHost := os.Getenv("MAILER_HOST")
	if mailerHost != "" {
		config.mailer.Host = mailerHost
	}
	mailerPortStr := os.Getenv("MAILER_PORT")
	if mailerPortStr != "" {
		p, err := strconv.Atoi(mailerPortStr)
		if err != nil {
			return fmt.Errorf("cannot parse mailer port from env variable : %w", err)
		}
		config.mailer.Port = p
	}
	mailerFromAddress := os.Getenv("MAILER_FROM_ADDRESS")
	if mailerFromAddress != "" {
		config.mailer.FromAddress = mailerFromAddress
	}
	mailerFromDisplayName := os.Getenv("MAILER_FROM_DISPLAY_NAME")
	if mailerFromDisplayName != "" {
		config.mailer.FromDisplayName = mailerFromDisplayName
	}
	mailerReplyTo := os.Getenv("MAILER_REPLY_TO")
	if mailerReplyTo != "" {
		config.mailer.ReplyTo = mailerReplyTo
	}

	// database
	databaseURI := os.Getenv("DATABASE_URI")
	if databaseURI != "" {
		config.database.URI = databaseURI
	}
	maxConnLifetimeStr := os.Getenv("DATABASE_MAX_CONN_LIFETIME")
	if maxConnLifetimeStr != "" {
		d, err := time.ParseDuration(maxConnLifetimeStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : max conn lifetime : %w", err)
		}
		config.database.MaxConnLifetime = d
	}
	maxConnLifetimeJitterStr := os.Getenv("DATABASE_MAX_CONN_LIFETIME_JITTER")
	if maxConnLifetimeJitterStr != "" {
		d, err := time.ParseDuration(maxConnLifetimeJitterStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : max conn lifetime jitter : %w", err)
		}
		config.database.MaxConnLifetimeJitter = d
	}
	maxConnIdleTimeStr := os.Getenv("DATABASE_MAX_CONN_IDEAL_TIME")
	if maxConnIdleTimeStr != "" {
		d, err := time.ParseDuration(maxConnIdleTimeStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : max conn ideal time : %w", err)
		}
		config.database.MaxConnIdleTime = d
	}
	healthCheckPeriodStr := os.Getenv("DATABASE_MAX_CONN_IDEAL_TIME")
	if healthCheckPeriodStr != "" {
		d, err := time.ParseDuration(healthCheckPeriodStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : health check period : %w", err)
		}
		config.database.HealthCheckPeriod = d
	}
	maxConnStr := os.Getenv("DATABASE_MAX_CONN")
	if maxConnStr != "" {
		c, err := strconv.Atoi(maxConnStr)
		if err != nil {
			return fmt.Errorf("cannot parse database max conn from env variable : %w", err)
		}
		config.database.MaxConns = int32(c)
	}
	minConnStr := os.Getenv("DATABASE_MIN_CONN")
	if minConnStr != "" {
		c, err := strconv.Atoi(minConnStr)
		if err != nil {
			return fmt.Errorf("cannot parse database min conn from env variable : %w", err)
		}
		config.database.MinConns = int32(c)
	}

	// storage
	storageEndpoint := os.Getenv("STORAGE_ENDPOINT")
	if storageEndpoint != "" {
		config.storage.Endpoint = storageEndpoint
	}
	storageAccessKey := os.Getenv("STORAGE_ACCESS_KEY")
	if storageAccessKey != "" {
		config.storage.AccessKey = storageAccessKey
	}
	storageSecret := os.Getenv("STORAGE_SECRET")
	if storageSecret != "" {
		config.storage.Secret = storageSecret
	}
	storageBucket := os.Getenv("STORAGE_BUCKET")
	if storageBucket != "" {
		config.storage.Bucket = storageBucket
	}
	storageBucketRegion := os.Getenv("STORAGE_BUCKET_REGION")
	if storageBucketRegion != "" {
		config.storage.Region = storageBucketRegion
	}
	storageUseSSL := os.Getenv("STORAGE_USE_SSL")
	if storageUseSSL == "true" {
		config.storage.UseSSL = true
	}
	storageAutoCreateBucket := os.Getenv("STORAGE_AUTO_CREATE_BUCKET")
	if storageAutoCreateBucket == "true" {
		config.storage.AutoCreateBucket = true
	}

	// auth
	accessTokenSecret := os.Getenv("AUTH_ACCESS_TOKEN_SECRET")
	if accessTokenSecret != "" {
		config.auth.AccessTokenSecret = accessTokenSecret
	}
	// TODO: rest
	// accessTokenLifetimeStr := os.Getenv("ACCESS_TOKEN_LIFETIME")
	// if accessTokenLifetime != "" {
	// 	config.auth.AccessTokenLifetime = accessTokenLifetime
	// }

	return nil
}

func (*config) parseFromCLIFlags() error {
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
	authConf := auth.DefaultConfig()
	return config{
		logLevel:      zerolog.InfoLevel,
		mailer:        &mailer.Config{},
		database:      &dbConf,
		appPublicHost: "",
		auth:          &authConf,
		storage:       &storage.Config{},
	}
}

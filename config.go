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
)

type Config struct {
	AppPublicHost string
	Mailer        *mailer.Config
	Database      *database.Config
	Storage       *storage.Config
	Auth          *auth.Config
}

func (config *Config) Validate() error {
	return validation.ValidateStruct(
		config,
		validation.Field(&config.AppPublicHost, validation.Required, is.URL),
		validation.Field(&config.Mailer),
		validation.Field(&config.Database),
		validation.Field(&config.Auth),
		validation.Field(&config.Storage),
	)
}

func (config *Config) ParseFromEnvVars() error {
	// host
	host := os.Getenv("APP_PUBLIC_HOST")
	if host != "" {
		config.AppPublicHost = host
		config.Auth.AppPublicHost = host
	}

	// mailer
	mailerUsername := os.Getenv("MAILER_USERNAME")
	if mailerUsername != "" {
		config.Mailer.Username = mailerUsername
	}
	mailerPassword := os.Getenv("MAILER_PASSWORD")
	if mailerPassword != "" {
		config.Mailer.Password = mailerPassword
	}
	mailerHost := os.Getenv("MAILER_HOST")
	if mailerHost != "" {
		config.Mailer.Host = mailerHost
	}
	mailerPortStr := os.Getenv("MAILER_PORT")
	if mailerPortStr != "" {
		p, err := strconv.Atoi(mailerPortStr)
		if err != nil {
			return fmt.Errorf("cannot parse mailer port from env variable : %w", err)
		}
		config.Mailer.Port = p
	}
	mailerFromAddress := os.Getenv("MAILER_FROM_ADDRESS")
	if mailerFromAddress != "" {
		config.Mailer.FromAddress = mailerFromAddress
	}
	mailerFromDisplayName := os.Getenv("MAILER_FROM_DISPLAY_NAME")
	if mailerFromDisplayName != "" {
		config.Mailer.FromDisplayName = mailerFromDisplayName
	}
	mailerReplyTo := os.Getenv("MAILER_REPLY_TO")
	if mailerReplyTo != "" {
		config.Mailer.ReplyTo = mailerReplyTo
	}

	// database
	databaseURI := os.Getenv("DATABASE_URI")
	if databaseURI != "" {
		config.Database.URI = databaseURI
	}
	maxConnLifetimeStr := os.Getenv("DATABASE_MAX_CONN_LIFETIME")
	if maxConnLifetimeStr != "" {
		d, err := time.ParseDuration(maxConnLifetimeStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : max conn lifetime : %w", err)
		}
		config.Database.MaxConnLifetime = d
	}
	maxConnLifetimeJitterStr := os.Getenv("DATABASE_MAX_CONN_LIFETIME_JITTER")
	if maxConnLifetimeJitterStr != "" {
		d, err := time.ParseDuration(maxConnLifetimeJitterStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : max conn lifetime jitter : %w", err)
		}
		config.Database.MaxConnLifetimeJitter = d
	}
	maxConnIdleTimeStr := os.Getenv("DATABASE_MAX_CONN_IDEAL_TIME")
	if maxConnIdleTimeStr != "" {
		d, err := time.ParseDuration(maxConnIdleTimeStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : max conn ideal time : %w", err)
		}
		config.Database.MaxConnIdleTime = d
	}
	healthCheckPeriodStr := os.Getenv("DATABASE_MAX_CONN_IDEAL_TIME")
	if healthCheckPeriodStr != "" {
		d, err := time.ParseDuration(healthCheckPeriodStr)
		if err != nil {
			return fmt.Errorf("cannot parse duration from env variable : health check period : %w", err)
		}
		config.Database.HealthCheckPeriod = d
	}
	maxConnStr := os.Getenv("DATABASE_MAX_CONN")
	if maxConnStr != "" {
		c, err := strconv.Atoi(maxConnStr)
		if err != nil {
			return fmt.Errorf("cannot parse database max conn from env variable : %w", err)
		}
		config.Database.MaxConns = int32(c)
	}
	minConnStr := os.Getenv("DATABASE_MIN_CONN")
	if minConnStr != "" {
		c, err := strconv.Atoi(minConnStr)
		if err != nil {
			return fmt.Errorf("cannot parse database min conn from env variable : %w", err)
		}
		config.Database.MinConns = int32(c)
	}

	// storage
	storageEndpoint := os.Getenv("STORAGE_ENDPOINT")
	if storageEndpoint != "" {
		config.Storage.Endpoint = storageEndpoint
	}
	storageAccessKey := os.Getenv("STORAGE_ACCESS_KEY")
	if storageAccessKey != "" {
		config.Storage.AccessKey = storageAccessKey
	}
	storageSecret := os.Getenv("STORAGE_SECRET")
	if storageSecret != "" {
		config.Storage.Secret = storageSecret
	}
	storageBucket := os.Getenv("STORAGE_BUCKET")
	if storageBucket != "" {
		config.Storage.Bucket = storageBucket
	}
	storageBucketRegion := os.Getenv("STORAGE_BUCKET_REGION")
	if storageBucketRegion != "" {
		config.Storage.Region = storageBucketRegion
	}
	storageUseSSL := os.Getenv("STORAGE_USE_SSL")
	if storageUseSSL == "true" {
		config.Storage.UseSSL = true
	}
	storageAutoCreateBucket := os.Getenv("STORAGE_AUTO_CREATE_BUCKET")
	if storageAutoCreateBucket == "true" {
		config.Storage.AutoCreateBucket = true
	}

	// auth
	accessTokenSecret := os.Getenv("AUTH_ACCESS_TOKEN_SECRET")
	if accessTokenSecret != "" {
		config.Auth.AccessTokenSecret = accessTokenSecret
	}
	// TODO: rest
	// accessTokenLifetimeStr := os.Getenv("ACCESS_TOKEN_LIFETIME")
	// if accessTokenLifetime != "" {
	// 	config.auth.AccessTokenLifetime = accessTokenLifetime
	// }

	return nil
}

func (*Config) ParseFromCLIFlags() error {
	return nil
}

func (config *Config) parse() error {
	if err := config.ParseFromEnvVars(); err != nil {
		return err
	}
	return config.ParseFromCLIFlags()
}

func DefaultConfig() Config {
	dbConf := database.DefaultConfig()
	authConf := auth.DefaultConfig()
	return Config{
		Mailer:        &mailer.Config{},
		Database:      &dbConf,
		AppPublicHost: "",
		Auth:          &authConf,
		Storage:       &storage.Config{},
	}
}

package bookmymovie

import (
	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/storage"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/rs/zerolog"
)

type Config struct {
	AppPublicHost string        `conf:"env:APP-PUBLIC_HOST,flag:app-public-host"`
	LogLevel      zerolog.Level `conf:"env:LOG_LEVEL,flag:log-level,default,INFO"`
	Mailer        *mailer.Config
	Database      *database.Config
	Storage       *storage.Config
	Auth          *auth.Config
}

func (config *Config) Validate() error {
	return validation.ValidateStruct(
		config,
		validation.Field(&config.AppPublicHost, validation.Required, is.URL),
		validation.Field(&config.LogLevel),
		validation.Field(&config.Mailer),
		validation.Field(&config.Database),
		validation.Field(&config.Auth),
		validation.Field(&config.Storage),
	)
}

// func (config *Config) parseFromEnvVars() error {
// 	// log level
// 	levelStr := os.Getenv("LOG_LEVEL")
// 	if levelStr == "" {
// 		levelStr = zerolog.InfoLevel.String()
// 	}

// 	// host
// 	host := os.Getenv("APP_HOST")
// 	if host != "" {
// 		config.AppPublicHost = host
// 		config.Auth.Host = host
// 	}

// 	logLevel, err := zerolog.ParseLevel(levelStr)
// 	if err != nil {
// 		return fmt.Errorf("invalid log level in env variable : %w", err)
// 	}
// 	config.LogLevel = logLevel

// 	// mailer
// 	mailerUsername := os.Getenv("SMTP_USERNAME")
// 	if mailerUsername != "" {
// 		config.Mailer.Username = mailerUsername
// 	}
// 	mailerPassword := os.Getenv("SMTP_PASSWORD")
// 	if mailerPassword != "" {
// 		config.Mailer.Password = mailerPassword
// 	}
// 	mailerHost := os.Getenv("SMTP_HOST")
// 	if mailerHost != "" {
// 		config.Mailer.Host = mailerHost
// 	}
// 	mailerPortStr := os.Getenv("SMTP_PORT")
// 	if mailerPortStr != "" {
// 		p, err := strconv.Atoi(mailerPortStr)
// 		if err != nil {
// 			return fmt.Errorf("cannot parse mailer port from env variable : %w", err)
// 		}
// 		config.Mailer.Port = p
// 	}
// 	mailerFromAddress := os.Getenv("SMTP_FROM_ADDRESS")
// 	if mailerFromAddress != "" {
// 		config.Mailer.FromAddress = mailerFromAddress
// 	}
// 	mailerFromDisplayName := os.Getenv("SMTP_FROM_DISPLAY_NAME")
// 	if mailerFromDisplayName != "" {
// 		config.Mailer.FromDisplayName = mailerFromDisplayName
// 	}
// 	mailerReplyTo := os.Getenv("SMTP_REPLY_TO")
// 	if mailerReplyTo != "" {
// 		config.Mailer.ReplyTo = mailerReplyTo
// 	}

// 	// database
// 	databaseURI := os.Getenv("DATABASE_URI")
// 	if databaseURI != "" {
// 		config.Database.URI = databaseURI
// 	}
// 	maxConnLifetimeStr := os.Getenv("DATABASE_MAX_CONN_LIFETIME")
// 	if maxConnLifetimeStr != "" {
// 		d, err := time.ParseDuration(maxConnLifetimeStr)
// 		if err != nil {
// 			return fmt.Errorf("cannot parse duration from env variable : max conn lifetime : %w", err)
// 		}
// 		config.Database.MaxConnLifetime = d
// 	}
// 	maxConnLifetimeJitterStr := os.Getenv("DATABASE_MAX_CONN_LIFETIME_JITTER")
// 	if maxConnLifetimeJitterStr != "" {
// 		d, err := time.ParseDuration(maxConnLifetimeJitterStr)
// 		if err != nil {
// 			return fmt.Errorf("cannot parse duration from env variable : max conn lifetime jitter : %w", err)
// 		}
// 		config.Database.MaxConnLifetimeJitter = d
// 	}
// 	maxConnIdleTimeStr := os.Getenv("DATABASE_MAX_CONN_IDEAL_TIME")
// 	if maxConnIdleTimeStr != "" {
// 		d, err := time.ParseDuration(maxConnIdleTimeStr)
// 		if err != nil {
// 			return fmt.Errorf("cannot parse duration from env variable : max conn ideal time : %w", err)
// 		}
// 		config.Database.MaxConnIdleTime = d
// 	}
// 	healthCheckPeriodStr := os.Getenv("DATABASE_MAX_CONN_IDEAL_TIME")
// 	if healthCheckPeriodStr != "" {
// 		d, err := time.ParseDuration(healthCheckPeriodStr)
// 		if err != nil {
// 			return fmt.Errorf("cannot parse duration from env variable : health check period : %w", err)
// 		}
// 		config.Database.HealthCheckPeriod = d
// 	}
// 	maxConnStr := os.Getenv("DATABASE_MAX_CONN")
// 	if maxConnStr != "" {
// 		c, err := strconv.Atoi(maxConnStr)
// 		if err != nil {
// 			return fmt.Errorf("cannot parse database max conn from env variable : %w", err)
// 		}
// 		config.Database.MaxConns = int32(c)
// 	}
// 	minConnStr := os.Getenv("DATABASE_MIN_CONN")
// 	if minConnStr != "" {
// 		c, err := strconv.Atoi(minConnStr)
// 		if err != nil {
// 			return fmt.Errorf("cannot parse database min conn from env variable : %w", err)
// 		}
// 		config.Database.MinConns = int32(c)
// 	}

// 	// storage
// 	storageEndpoint := os.Getenv("S3_STORAGE_ENDPOINT")
// 	if storageEndpoint != "" {
// 		config.Storage.Endpoint = storageEndpoint
// 	}
// 	storageAccessKey := os.Getenv("S3_STORAGE_ACCESS_KEY")
// 	if storageAccessKey != "" {
// 		config.Storage.AccessKey = storageAccessKey
// 	}
// 	storageSecret := os.Getenv("S3_STORAGE_SECRET")
// 	if storageSecret != "" {
// 		config.Storage.Secret = storageSecret
// 	}
// 	storageBucket := os.Getenv("S3_STORAGE_BUCKET")
// 	if storageBucket != "" {
// 		config.Storage.Bucket = storageBucket
// 	}
// 	storageBucketRegion := os.Getenv("S3_STORAGE_BUCKET_REGION")
// 	if storageBucketRegion != "" {
// 		config.Storage.Region = storageBucketRegion
// 	}
// 	storageUseSSL := os.Getenv("S3_STORAGE_USE_SSL")
// 	if storageUseSSL == "true" {
// 		config.Storage.UseSSL = true
// 	}
// 	storageAutoCreateBucket := os.Getenv("S3_STORAGE_AUTO_CREATE_BUCKET")
// 	if storageAutoCreateBucket == "true" {
// 		config.Storage.AutoCreateBucket = true
// 	}

// 	// auth
// 	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
// 	if accessTokenSecret != "" {
// 		config.Auth.AccessTokenSecret = accessTokenSecret
// 	}
// 	// TODO: rest
// 	// accessTokenLifetimeStr := os.Getenv("ACCESS_TOKEN_LIFETIME")
// 	// if accessTokenLifetime != "" {
// 	// 	config.auth.AccessTokenLifetime = accessTokenLifetime
// 	// }

// 	return nil
// }

// func (*Config) parseFromCLIFlags() error {
// 	return nil
// }

// func (config *Config) parse() error {
// 	if err := config.parseFromEnvVars(); err != nil {
// 		return err
// 	}
// 	return config.parseFromCLIFlags()
// }

func DefaultConfig() Config {
	dbConf := database.DefaultConfig()
	authConf := auth.DefaultConfig()
	return Config{
		LogLevel:      zerolog.InfoLevel,
		Mailer:        &mailer.Config{},
		Database:      &dbConf,
		AppPublicHost: "",
		Auth:          &authConf,
		Storage:       &storage.Config{},
	}
}

package auth

import (
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"github.com/rs/zerolog"
)

type Auth struct {
	config *Config
	logger *zerolog.Logger
	db     *database.Database
	mailer *mailer.Mailer

	revokedTokens map[int64]time.Time
}

func New(config *Config, logger *zerolog.Logger, db *database.Database, m *mailer.Mailer) Auth {
	if err := config.Validate(); err != nil {
		logger.Fatal().Err(err).Msg("auth config validation failed")
	}
	a := Auth{
		logger:        logger,
		db:            db,
		mailer:        m,
		config:        config,
		revokedTokens: make(map[int64]time.Time),
	}
	a.startBackgroundRevokedTokenCleanup()

	return a
}

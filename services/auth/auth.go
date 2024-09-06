package auth

import (
	"errors"
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

func New(config *Config, logger *zerolog.Logger, db *database.Database, m *mailer.Mailer) (Auth, error) {
	if err := config.Validate(); err != nil {
		return Auth{}, errors.Join(errors.New("auth config validation failed"), err)
	}
	a := Auth{
		logger:        logger,
		db:            db,
		mailer:        m,
		config:        config,
		revokedTokens: make(map[int64]time.Time),
	}
	a.startBackgroundRevokedTokenCleanup()

	return a, nil
}

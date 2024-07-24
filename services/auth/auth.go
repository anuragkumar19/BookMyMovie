package auth

import (
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"github.com/rs/zerolog"
)

type Auth struct {
	config *AuthConfig
	logger *zerolog.Logger
	db     *database.Database
	mailer *mailer.Mailer

	revokedTokens map[int64]time.Time
}

func New(config *AuthConfig, logger *zerolog.Logger, db *database.Database, mailer *mailer.Mailer) Auth {
	a := Auth{
		logger:        logger,
		db:            db,
		mailer:        mailer,
		config:        config,
		revokedTokens: make(map[int64]time.Time),
	}
	a.startBackgroundRevokedTokenCleanup()

	return a
}

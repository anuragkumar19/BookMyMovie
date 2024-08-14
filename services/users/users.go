package users

import (
	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"bookmymovie.app/bookmymovie/services/auth"
	"github.com/rs/zerolog"
)

type Users struct {
	logger *zerolog.Logger
	db     *database.Database
	mailer *mailer.Mailer

	auth *auth.Auth
}

func New(logger *zerolog.Logger, db *database.Database, m *mailer.Mailer, authService *auth.Auth) Users {
	return Users{
		logger: logger,
		db:     db,
		mailer: m,
		auth:   authService,
	}
}

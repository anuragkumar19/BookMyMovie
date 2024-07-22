package auth

import (
	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"github.com/rs/zerolog"
)

type Auth struct {
	logger *zerolog.Logger
	db     *database.Database
	mailer *mailer.Mailer
}

func New(logger *zerolog.Logger, db *database.Database, mailer *mailer.Mailer) Auth {
	return Auth{
		logger: logger,
		db:     db,
		mailer: mailer,
	}
}

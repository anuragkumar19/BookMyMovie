package movies

import (
	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"github.com/rs/zerolog"
)

type Movies struct {
	logger *zerolog.Logger
	db     *database.Database
	auth   *auth.Auth
}

func New(logger *zerolog.Logger, db *database.Database, auth *auth.Auth) Movies {
	return Movies{
		logger: logger,
		db:     db,
		auth:   auth,
	}
}

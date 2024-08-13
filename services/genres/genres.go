package genres

import (
	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"github.com/rs/zerolog"
)

type Genres struct {
	logger *zerolog.Logger
	db     *database.Database
	auth   *auth.Auth
}

func New(logger *zerolog.Logger, db *database.Database, auth *auth.Auth) Genres {
	return Genres{
		logger: logger,
		db:     db,
		auth:   auth,
	}
}

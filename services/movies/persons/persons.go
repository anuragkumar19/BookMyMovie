package persons

import (
	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/storage"
	"github.com/rs/zerolog"
)

type Persons struct {
	logger  *zerolog.Logger
	db      *database.Database
	storage *storage.Storage
	auth    *auth.Auth
}

func New(logger *zerolog.Logger, db *database.Database, s *storage.Storage, a *auth.Auth) Persons {
	return Persons{
		logger:  logger,
		db:      db,
		storage: s,
		auth:    a,
	}
}

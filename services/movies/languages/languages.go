package languages

import (
	"context"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"github.com/rs/zerolog"
)

type Languages struct {
	logger *zerolog.Logger
	db     *database.Database
	auth   *auth.Auth

	cache *cache
}

func New(logger *zerolog.Logger, db *database.Database, auth *auth.Auth) Languages {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	languages, err := db.GetAllMoviesLanguages(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load languages from database")
	}
	return Languages{
		logger: logger,
		db:     db,
		auth:   auth,
		cache: &cache{
			languages: languages,
		},
	}
}

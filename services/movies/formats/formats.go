package formats

import (
	"context"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"github.com/rs/zerolog"
)

type Formats struct {
	logger *zerolog.Logger
	db     *database.Database
	auth   *auth.Auth

	cache *cache
}

func New(logger *zerolog.Logger, db *database.Database, auth *auth.Auth) Formats {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	formats, err := db.GetAllMoviesFormats(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load formats from database")
	}
	return Formats{
		logger: logger,
		db:     db,
		auth:   auth,
		cache: &cache{
			formats: formats,
		},
	}
}

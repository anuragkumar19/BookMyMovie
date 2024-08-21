package genres

import (
	"context"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"github.com/rs/zerolog"
)

type Genres struct {
	logger *zerolog.Logger
	db     *database.Database
	auth   *auth.Auth

	cache *cache
}

func New(logger *zerolog.Logger, db *database.Database, auth *auth.Auth) Genres {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	genres, err := db.GetAllMoviesGenres(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load genres from database")
	}
	return Genres{
		logger: logger,
		db:     db,
		auth:   auth,
		cache: &cache{
			genres: genres,
		},
	}
}

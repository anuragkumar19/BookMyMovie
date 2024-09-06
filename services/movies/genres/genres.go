package genres

import (
	"context"
	"errors"
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

func New(logger *zerolog.Logger, db *database.Database, a *auth.Auth) (Genres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	genres, err := db.GetAllMoviesGenres(ctx)
	if err != nil {
		return Genres{}, errors.Join(errors.New("failed to load genres from database"), err)
	}
	return Genres{
		logger: logger,
		db:     db,
		auth:   a,
		cache: &cache{
			genres: genres,
		},
	}, nil
}

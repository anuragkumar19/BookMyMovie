package genres

import (
	"context"
	"errors"

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

func New(ctx context.Context, logger *zerolog.Logger, db *database.Database, a *auth.Auth) (Genres, error) {
	// init cache
	c := &cache{}
	genres, err := db.GetAllMoviesGenres(ctx)
	if err != nil {
		return Genres{}, errors.Join(errors.New("failed to load genres from database"), err)
	}
	c.refresh(genres)

	return Genres{
		logger: logger,
		db:     db,
		auth:   a,
		cache:  c,
	}, nil
}

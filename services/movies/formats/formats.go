package formats

import (
	"context"
	"errors"
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

func New(logger *zerolog.Logger, db *database.Database, a *auth.Auth) (Formats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	formats, err := db.GetAllMoviesFormats(ctx)
	if err != nil {
		return Formats{}, errors.Join(errors.New("failed to load formats from database"), err)
	}
	return Formats{
		logger: logger,
		db:     db,
		auth:   a,
		cache: &cache{
			formats: formats,
		},
	}, nil
}

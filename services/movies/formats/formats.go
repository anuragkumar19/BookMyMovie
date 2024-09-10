package formats

import (
	"context"
	"errors"

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

func New(ctx context.Context, logger *zerolog.Logger, db *database.Database, a *auth.Auth) (Formats, error) {
	// init cache
	c := &cache{}
	formats, err := db.GetAllMoviesFormats(ctx)
	if err != nil {
		return Formats{}, errors.Join(errors.New("failed to load formats from database"), err)
	}
	c.refresh(formats)

	return Formats{
		logger: logger,
		db:     db,
		auth:   a,
		cache:  c,
	}, nil
}

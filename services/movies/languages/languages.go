package languages

import (
	"context"
	"errors"

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

func New(ctx context.Context, logger *zerolog.Logger, db *database.Database, a *auth.Auth) (Languages, error) {
	// init cache
	c := &cache{}
	languages, err := db.GetAllMoviesLanguages(ctx)
	if err != nil {
		return Languages{}, errors.Join(errors.New("failed to load languages from database"), err)
	}
	c.refresh(languages)

	return Languages{
		logger: logger,
		db:     db,
		auth:   a,
		cache:  c,
	}, nil
}

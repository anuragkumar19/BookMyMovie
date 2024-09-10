package languages

import (
	"context"
	"errors"
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

func New(logger *zerolog.Logger, db *database.Database, a *auth.Auth) (Languages, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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

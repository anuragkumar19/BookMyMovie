package languages

import (
	"context"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	services_errors "bookmymovie.app/bookmymovie/services/errors"
)

func (s *Languages) GetByID(_ context.Context, id string) (database.MoviesLanguage, error) {
	i := slices.IndexFunc(s.cache.languages, func(l database.MoviesLanguage) bool {
		return l.ID == id
	})
	if i == -1 {
		return database.MoviesLanguage{}, services_errors.ErrNotFound
	}
	return s.cache.languages[i], nil
}

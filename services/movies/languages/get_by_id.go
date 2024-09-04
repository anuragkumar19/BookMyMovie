package languages

import (
	"context"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	serviceserrorss "bookmymovie.app/bookmymovie/services/serviceserrors"
)

func (s *Languages) GetByID(_ context.Context, id string) (database.MoviesLanguage, error) {
	i := slices.IndexFunc(s.cache.languages, func(l database.MoviesLanguage) bool {
		return l.ID == id
	})
	if i == -1 {
		return database.MoviesLanguage{}, serviceserrorss.ErrNotFound
	}
	return s.cache.languages[i], nil
}

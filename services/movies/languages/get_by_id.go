package languages

import (
	"context"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
)

func (s *Languages) GetByID(_ context.Context, id int64) (database.MoviesLanguage, error) {
	i := slices.IndexFunc(s.cache.languages, func(l database.MoviesLanguage) bool {
		return l.ID == id
	})
	if i == -1 {
		return database.MoviesLanguage{}, serviceserrors.New(serviceserrors.ErrorTypeNotFound, "movie language not found")
	}
	return s.cache.languages[i], nil
}

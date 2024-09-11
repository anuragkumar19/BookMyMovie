package languages

import (
	"context"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
)

func (s *Languages) GetByID(_ context.Context, id int64) (database.MoviesLanguage, error) {
	l, ok := s.cache.index[id]
	if !ok {
		return database.MoviesLanguage{}, services.NewError(services.ErrorTypeNotFound, "movies language not found")
	}
	return l, nil
}

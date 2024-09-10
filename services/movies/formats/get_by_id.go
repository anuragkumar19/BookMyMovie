package formats

import (
	"context"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
)

func (s *Formats) GetByID(_ context.Context, id int64) (database.MoviesFormat, error) {
	l, ok := s.cache.index[id]
	if !ok {
		return database.MoviesFormat{}, serviceserrors.New(serviceserrors.ErrorTypeNotFound, "movies format not found")
	}
	return l, nil
}

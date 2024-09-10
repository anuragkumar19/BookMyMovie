package genres

import (
	"context"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
)

func (s *Genres) GetByID(_ context.Context, id int64) (database.MoviesGenre, error) {
	l, ok := s.cache.index[id]
	if !ok {
		return database.MoviesGenre{}, serviceserrors.New(serviceserrors.ErrorTypeNotFound, "movies genre not found")
	}
	return l, nil
}

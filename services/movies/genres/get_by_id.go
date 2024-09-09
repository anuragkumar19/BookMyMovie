package genres

import (
	"context"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
)

func (s *Genres) GetByID(_ context.Context, id int64) (database.MoviesGenre, error) {
	i := slices.IndexFunc(s.cache.genres, func(g database.MoviesGenre) bool {
		return g.ID == id
	})
	if i == -1 {
		return database.MoviesGenre{}, serviceserrors.New(serviceserrors.ErrorTypeNotFound, "movie genre not found")
	}
	return s.cache.genres[i], nil
}

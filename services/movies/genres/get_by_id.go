package genres

import (
	"context"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	services_errors "bookmymovie.app/bookmymovie/services/errors"
)

func (s *Genres) GetByID(_ context.Context, id string) (database.MoviesGenre, error) {
	i := slices.IndexFunc(s.cache.genres, func(g database.MoviesGenre) bool {
		return g.ID == id
	})
	if i == -1 {
		return database.MoviesGenre{}, services_errors.ErrNotFound
	}
	return s.cache.genres[i], nil
}

package genres

import (
	"context"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
)

func (s *Genres) GetByID(_ context.Context, id string) (database.MoviesGenre, error) {
	i := slices.IndexFunc(s.cache.genres, func(g database.MoviesGenre) bool {
		return g.ID == id
	})
	if i == -1 {
		return database.MoviesGenre{}, serviceserrors.ErrNotFound
	}
	return s.cache.genres[i], nil
}

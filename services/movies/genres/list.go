package genres

import (
	"context"

	"bookmymovie.app/bookmymovie/database"
)

func (s *Genres) List(_ context.Context) ([]database.MoviesGenre, error) {
	return s.cache.genres, nil
}

package formats

import (
	"context"

	"bookmymovie.app/bookmymovie/database"
)

func (s *Formats) List(_ context.Context) ([]database.MoviesFormat, error) {
	return s.cache.formats, nil
}

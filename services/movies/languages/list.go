package languages

import (
	"context"

	"bookmymovie.app/bookmymovie/database"
)

func (s *Languages) List(_ context.Context) ([]database.MoviesLanguage, error) {
	return s.cache.languages, nil
}

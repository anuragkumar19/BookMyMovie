package formats

import (
	"context"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	services_errors "bookmymovie.app/bookmymovie/services/errors"
)

func (s *Formats) GetByID(_ context.Context, id string) (database.MoviesFormat, error) {
	i := slices.IndexFunc(s.cache.formats, func(f database.MoviesFormat) bool {
		return f.ID == id
	})
	if i == -1 {
		return database.MoviesFormat{}, services_errors.ErrNotFound
	}
	return s.cache.formats[i], nil
}
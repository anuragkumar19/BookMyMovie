package formats

import (
	"context"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
)

func (s *Formats) GetByID(_ context.Context, id int64) (database.MoviesFormat, error) {
	i := slices.IndexFunc(s.cache.formats, func(f database.MoviesFormat) bool {
		return f.ID == id
	})
	if i == -1 {
		return database.MoviesFormat{}, serviceserrors.ErrNotFound
	}
	return s.cache.formats[i], nil
}

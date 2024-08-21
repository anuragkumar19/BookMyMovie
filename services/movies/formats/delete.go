package formats

import (
	"context"
	"errors"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	services_errors "bookmymovie.app/bookmymovie/services/errors"
	"github.com/jackc/pgx/v5"
)

func (s *Formats) Delete(ctx context.Context, accessToken string, id string) error {
	authMetadata, err := s.auth.GetAuthMetadata(accessToken)
	if err != nil {
		return err
	}
	if err := s.auth.CheckPermissions(&authMetadata, auth.MoviesFormatsDelete); err != nil {
		return err
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		if errors.Is(err, services_errors.ErrNotFound) {
			return services_errors.ErrNotFound
		}
		return err
	}

	if err := s.db.DeleteMoviesFormat(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return services_errors.ErrUpdateConflict
		}
		return err
	}

	formats := make([]database.MoviesFormat, len(s.cache.formats))
	copy(formats, s.cache.formats)

	formats = slices.DeleteFunc(formats, func(f database.MoviesFormat) bool {
		return f.ID == id
	})
	s.cache.refresh(formats)
	return nil
}

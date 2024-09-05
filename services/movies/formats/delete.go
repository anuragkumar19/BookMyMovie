package formats

import (
	"context"
	"errors"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
	"github.com/jackc/pgx/v5"
)

func (s *Formats) Delete(ctx context.Context, authMeta *auth.Metadata, id string) error {
	if err := authMeta.Valid(); err != nil {
		return err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesFormatsDelete); err != nil {
		return err
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		if errors.Is(err, serviceserrors.ErrNotFound) {
			return serviceserrors.ErrNotFound
		}
		return err
	}

	if err := s.db.DeleteMoviesFormat(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return serviceserrors.ErrUpdateConflict
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

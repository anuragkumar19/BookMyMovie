package languages

import (
	"context"
	"errors"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	serviceserrorss "bookmymovie.app/bookmymovie/services/serviceserrors"
	"github.com/jackc/pgx/v5"
)

func (s *Languages) Delete(ctx context.Context, accessToken string, id string) error {
	authMetadata, err := s.auth.GetAuthMetadata(accessToken)
	if err != nil {
		return err
	}
	if err := s.auth.CheckPermissions(&authMetadata, auth.MoviesLanguagesDelete); err != nil {
		return err
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		if errors.Is(err, serviceserrorss.ErrNotFound) {
			return serviceserrorss.ErrNotFound
		}
		return err
	}

	if err := s.db.DeleteMoviesLanguage(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return serviceserrorss.ErrUpdateConflict
		}
		return err
	}

	langs := make([]database.MoviesLanguage, len(s.cache.languages))
	copy(langs, s.cache.languages)

	langs = slices.DeleteFunc(langs, func(l database.MoviesLanguage) bool {
		return l.ID == id
	})
	s.cache.refresh(langs)
	return nil
}

package languages

import (
	"context"
	"errors"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
	"github.com/jackc/pgx/v5"
)

func (s *Languages) Delete(ctx context.Context, authMeta *auth.Metadata, id int64) error {
	if err := authMeta.Valid(); err != nil {
		return err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesLanguagesDelete); err != nil {
		return err
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		return err
	}

	exist, err := s.db.CheckIfAnyMoviesAvailableAudioLanguagesExist(ctx, id)
	if err != nil {
		return err
	}
	if exist {
		return serviceserrors.New(serviceserrors.ErrorTypeInvalidArgument, "movies language is linked to one or many movies, so it cannot be deleted")
	}

	exist2, err := s.db.CheckIfAnyMoviesAvailableSubtitleLanguagesExist(ctx, id)
	if err != nil {
		return err
	}
	if exist2 {
		return serviceserrors.New(serviceserrors.ErrorTypeInvalidArgument, "movies language is linked to one or many movies, so it cannot be deleted")
	}

	if err := s.db.DeleteMoviesLanguage(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return serviceserrors.New(serviceserrors.ErrorConflict, "")
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

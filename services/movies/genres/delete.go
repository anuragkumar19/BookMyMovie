package genres

import (
	"context"
	"errors"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
	"github.com/jackc/pgx/v5"
)

func (s *Genres) Delete(ctx context.Context, authMeta *auth.Metadata, id string) error {
	if err := authMeta.Valid(); err != nil {
		return err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesGenresDelete); err != nil {
		return err
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		if errors.Is(err, serviceserrors.ErrNotFound) {
			return serviceserrors.ErrNotFound
		}
		return err
	}

	// TODO: check if any movies are linked to it

	if err := s.db.DeleteMoviesGenre(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return serviceserrors.ErrUpdateConflict
		}
		return err
	}

	genres := make([]database.MoviesGenre, len(s.cache.genres))
	copy(genres, s.cache.genres)

	genres = slices.DeleteFunc(genres, func(g database.MoviesGenre) bool {
		return g.ID == id
	})
	s.cache.refresh(genres)
	return nil
}

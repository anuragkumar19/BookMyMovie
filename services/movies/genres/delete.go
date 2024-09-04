package genres

import (
	"context"
	"errors"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	services_errors "bookmymovie.app/bookmymovie/services/serviceserrors"
	"github.com/jackc/pgx/v5"
)

func (s *Genres) Delete(ctx context.Context, accessToken string, id string) error {
	authMetadata, err := s.auth.GetAuthMetadata(accessToken)
	if err != nil {
		return err
	}
	if err := s.auth.CheckPermissions(&authMetadata, auth.MoviesGenresDelete); err != nil {
		return err
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		if errors.Is(err, services_errors.ErrNotFound) {
			return services_errors.ErrNotFound
		}
		return err
	}

	if err := s.db.DeleteMoviesGenre(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return services_errors.ErrUpdateConflict
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

package genres

import (
	"context"
	"errors"
	"slices"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	"bookmymovie.app/bookmymovie/services/auth"
	"github.com/jackc/pgx/v5"
)

func (s *Genres) Delete(ctx context.Context, authMeta *auth.Metadata, id int64) error {
	if err := authMeta.Valid(); err != nil {
		return err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesGenresDelete); err != nil {
		return err
	}

	if _, err := s.GetByID(ctx, id); err != nil {
		return err
	}

	exist, err := s.db.CheckIfAnyMoviesGenresJoinExist(ctx, id)
	if err != nil {
		return err
	}
	if exist {
		return services.NewError(services.ErrorTypeInvalidArgument, "movies genre is linked to one or many movies, so it cannot be deleted")
	}

	if err := s.db.DeleteMoviesGenre(ctx, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return services.NewError(services.ErrorTypeConflict, "")
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

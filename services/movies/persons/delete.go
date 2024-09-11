package persons

import (
	"context"
	"errors"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	"bookmymovie.app/bookmymovie/services/auth"
	"github.com/jackc/pgx/v5"
)

func (s *Persons) Delete(ctx context.Context, authMeta *auth.Metadata, id int64) error {
	if err := authMeta.Valid(); err != nil {
		return err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesPersonsDelete); err != nil {
		return err
	}

	person, err := s.db.GetMoviesPerson(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return services.NewError(services.ErrorTypeNotFound, "movies person not found")
		}
	}

	if err := s.db.DeleteMoviesPerson(ctx, &database.DeleteMoviesPersonParams{
		ID:      person.ID,
		Version: person.Version,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return services.NewError(services.ErrorTypeConflict, "")
		}
		return err
	}

	return nil
}

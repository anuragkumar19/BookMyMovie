package persons

import (
	"context"
	"errors"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	"github.com/jackc/pgx/v5"
)

func (s *Persons) GetById(ctx context.Context, id int64) (database.GetMoviesPersonRow, error) {
	person, err := s.db.GetMoviesPerson(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return database.GetMoviesPersonRow{}, services.NewError(services.ErrorTypeNotFound, "movies person not found")
		}
	}

	return person, nil
}

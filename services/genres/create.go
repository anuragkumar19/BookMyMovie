package genres

import (
	"context"
	"errors"
	"strings"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5"
)

type CreateGenreParams struct {
	AccessToken string
	ID          string
	DisplayName string
	About       string
}

func (data *CreateGenreParams) Transform() *CreateGenreParams {
	data.ID = strings.TrimSpace(data.ID)
	data.DisplayName = strings.TrimSpace(data.DisplayName)
	data.About = strings.TrimSpace(data.About)
	return data
}

func (data *CreateGenreParams) Validate() error {
	return validation.ValidateStruct(
		data,
		validation.Field(&data.ID, validation.Required),
		validation.Field(&data.DisplayName, validation.Required),
		validation.Field(&data.About, validation.Required),
	)
}

func (s *Genres) CreateGenre(ctx context.Context, params *CreateGenreParams) (database.MovieGenre, error) {
	authMeta, err := s.auth.GetAuthMetadata(params.AccessToken)
	if err != nil {
		return database.MovieGenre{}, err
	}
	if err := s.auth.CheckPermissions(&authMeta, auth.GenreCreate); err != nil {
		return database.MovieGenre{}, err
	}

	if _, err := s.db.GetGenreByID(ctx, params.ID); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return database.MovieGenre{}, err
	}

	genre, err := s.db.CreateGenre(ctx, &database.CreateGenreParams{
		ID:          params.ID,
		DisplayName: params.DisplayName,
		About:       params.About,
	})
	if err != nil {
		return database.MovieGenre{}, err
	}
	return genre, nil
}

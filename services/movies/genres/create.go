package genres

import (
	"context"
	"errors"
	"strings"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	services_errors "bookmymovie.app/bookmymovie/services/errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gosimple/slug"
)

type CreateParams struct {
	AccessToken string
	DisplayName string
	About       string
}

func (params *CreateParams) Transform() *CreateParams {
	params.DisplayName = strings.TrimSpace(params.DisplayName)
	params.About = strings.TrimSpace(params.About)
	return params
}

func (params *CreateParams) Validate() error {
	return validation.ValidateStruct(params,
		validation.Field(&params.DisplayName, validation.Required),
		validation.Field(&params.About, validation.Required),
	)
}

func (s *Genres) Create(ctx context.Context, params *CreateParams) (id string, err error) {
	authMetadata, err := s.auth.GetAuthMetadata(params.AccessToken)
	if err != nil {
		return "", err
	}
	if err := s.auth.CheckPermissions(&authMetadata, auth.GenreCreate); err != nil {
		return "", err
	}

	id = slug.Make(params.DisplayName)

	exist := true
	if _, err := s.GetByID(ctx, id); err != nil {
		if !errors.Is(err, services_errors.ErrNotFound) {
			return "", err
		}
		exist = false
	}
	if exist {
		return "", services_errors.ErrAlreadyExist
	}

	genre, err := s.db.CreateMoviesGenre(ctx, &database.CreateMoviesGenreParams{
		ID:          id,
		DisplayName: params.DisplayName,
		About:       params.About,
	})
	if err != nil {
		return "", err
	}

	genres := make([]database.MoviesGenre, len(s.cache.genres))
	copy(genres, s.cache.genres)

	genres = append(genres, genre)
	s.cache.refresh(genres)
	return id, nil
}
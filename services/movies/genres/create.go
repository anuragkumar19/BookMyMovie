package genres

import (
	"context"
	"errors"
	"strings"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gosimple/slug"
)

type CreateParams struct {
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

func (s *Genres) Create(ctx context.Context, authMeta *auth.Metadata, params *CreateParams) (id string, err error) {
	if err := authMeta.Valid(); err != nil {
		return "", err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesGenresCreate); err != nil {
		return "", err
	}

	id = slug.Make(params.DisplayName)

	exist := true
	if _, err := s.GetByID(ctx, id); err != nil {
		if !errors.Is(err, serviceserrors.ErrNotFound) {
			return "", err
		}
		exist = false
	}
	if exist {
		return "", serviceserrors.ErrAlreadyExist
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

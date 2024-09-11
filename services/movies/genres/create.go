package genres

import (
	"context"
	"strings"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	"bookmymovie.app/bookmymovie/services/auth"
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

func (params CreateParams) Validate() error {
	return validation.ValidateStruct(&params,
		validation.Field(&params.DisplayName, validation.Required),
		validation.Field(&params.About, validation.Required),
	)
}

func (s *Genres) Create(ctx context.Context, authMeta *auth.Metadata, params *CreateParams) (database.MoviesGenre, error) {
	if err := authMeta.Valid(); err != nil {
		return database.MoviesGenre{}, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesGenresCreate); err != nil {
		return database.MoviesGenre{}, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return database.MoviesGenre{}, err
		}
		return database.MoviesGenre{}, services.NewError(services.ErrorTypeInvalidArgument, err.Error())
	}

	slg := slug.Make(params.DisplayName)

	genre, err := s.db.CreateMoviesGenre(ctx, &database.CreateMoviesGenreParams{
		Slug:        slg,
		DisplayName: params.DisplayName,
		About:       params.About,
	})
	if err != nil {
		return database.MoviesGenre{}, err
	}

	genres := make([]database.MoviesGenre, len(s.cache.genres)+1)
	copy(genres, s.cache.genres)

	genres = append(genres, genre)
	s.cache.refresh(genres)
	return genre, nil
}

package languages

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
	EnglishName string
}

func (params *CreateParams) Transform() *CreateParams {
	params.DisplayName = strings.TrimSpace(params.DisplayName)
	params.EnglishName = strings.TrimSpace(params.EnglishName)
	return params
}

func (params CreateParams) Validate() error {
	return validation.ValidateStruct(&params,
		validation.Field(&params.DisplayName, validation.Required),
		validation.Field(&params.EnglishName, validation.Required),
	)
}

func (s *Languages) Create(ctx context.Context, authMeta *auth.Metadata, params *CreateParams) (database.MoviesLanguage, error) {
	if err := authMeta.Valid(); err != nil {
		return database.MoviesLanguage{}, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesLanguagesCreate); err != nil {
		return database.MoviesLanguage{}, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return database.MoviesLanguage{}, err
		}
		return database.MoviesLanguage{}, services.NewError(services.ErrorTypeInvalidArgument, err.Error())
	}

	slg := slug.Make(params.DisplayName)

	lang, err := s.db.CreateMoviesLanguage(ctx, &database.CreateMoviesLanguageParams{
		Slug:        slg,
		DisplayName: params.DisplayName,
		EnglishName: params.EnglishName,
	})
	if err != nil {
		return database.MoviesLanguage{}, err
	}

	langs := make([]database.MoviesLanguage, len(s.cache.languages)+1)
	copy(langs, s.cache.languages)

	langs = append(langs, lang)
	s.cache.refresh(langs)
	return lang, nil
}

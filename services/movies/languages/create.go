package languages

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
	EnglishName string
}

func (params *CreateParams) Transform() *CreateParams {
	params.DisplayName = strings.TrimSpace(params.DisplayName)
	params.EnglishName = strings.TrimSpace(params.EnglishName)
	return params
}

func (params *CreateParams) Validate() error {
	return validation.ValidateStruct(params,
		validation.Field(&params.DisplayName, validation.Required),
		validation.Field(&params.EnglishName, validation.Required),
	)
}

func (s *Languages) Create(ctx context.Context, authMeta *auth.Metadata, params *CreateParams) (id string, err error) {
	if err := authMeta.Valid(); err != nil {
		return "", err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesLanguagesCreate); err != nil {
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

	lang, err := s.db.CreateMoviesLanguage(ctx, &database.CreateMoviesLanguageParams{
		ID:          id,
		DisplayName: params.DisplayName,
		EnglishName: params.EnglishName,
	})
	if err != nil {
		return "", err
	}

	langs := make([]database.MoviesLanguage, len(s.cache.languages))
	copy(langs, s.cache.languages)

	langs = append(langs, lang)
	s.cache.refresh(langs)
	return id, nil
}

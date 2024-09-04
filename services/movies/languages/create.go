package languages

import (
	"context"
	"errors"
	"strings"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	serviceserrorss "bookmymovie.app/bookmymovie/services/serviceserrors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gosimple/slug"
)

type CreateParams struct {
	AccessToken string
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

func (s *Languages) Create(ctx context.Context, params *CreateParams) (id string, err error) {
	authMetadata, err := s.auth.GetAuthMetadata(params.AccessToken)
	if err != nil {
		return "", err
	}
	if err := s.auth.CheckPermissions(&authMetadata, auth.MoviesLanguagesCreate); err != nil {
		return "", err
	}

	id = slug.Make(params.DisplayName)

	exist := true
	if _, err := s.GetByID(ctx, id); err != nil {
		if !errors.Is(err, serviceserrorss.ErrNotFound) {
			return "", err
		}
		exist = false
	}
	if exist {
		return "", serviceserrorss.ErrAlreadyExist
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

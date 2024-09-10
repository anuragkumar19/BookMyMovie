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
	"github.com/jackc/pgx/v5"
)

type UpdateParams struct {
	ID          int64
	DisplayName *string
	EnglishName *string
}

func (params *UpdateParams) Transform() *UpdateParams {
	if params.DisplayName != nil {
		dn := strings.TrimSpace(*params.DisplayName)
		params.DisplayName = &dn
	}
	if params.EnglishName != nil {
		enm := strings.TrimSpace(*params.EnglishName)
		params.EnglishName = &enm
	}
	return params
}

func (params *UpdateParams) Validate() error {
	return validation.ValidateStruct(params,
		validation.Field(&params.DisplayName, validation.Length(1, 0)),
		validation.Field(&params.EnglishName, validation.Length(1, 0)),
	)
}

func (s *Languages) Update(ctx context.Context, authMeta *auth.Metadata, params *UpdateParams) (database.MoviesLanguage, error) {
	if err := authMeta.Valid(); err != nil {
		return database.MoviesLanguage{}, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesLanguagesUpdate); err != nil {
		return database.MoviesLanguage{}, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return database.MoviesLanguage{}, err
		}
		return database.MoviesLanguage{}, serviceserrors.New(serviceserrors.ErrorTypeInvalidArgument, err.Error())
	}

	language, err := s.GetByID(ctx, params.ID)
	if err != nil {
		return database.MoviesLanguage{}, nil
	}

	if params.DisplayName != nil {
		language.DisplayName = *params.DisplayName
	}
	if params.EnglishName != nil {
		language.EnglishName = *params.EnglishName
	}
	language.Slug = slug.Make(language.DisplayName)
	if err := s.db.UpdateMoviesLanguage(ctx, &database.UpdateMoviesLanguageParams{
		Slug:        language.Slug,
		DisplayName: language.DisplayName,
		EnglishName: language.EnglishName,
		ID:          language.ID,
		Version:     language.Version,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return database.MoviesLanguage{}, serviceserrors.New(serviceserrors.ErrorConflict, "")
		}
		return database.MoviesLanguage{}, err
	}

	return language, nil
}

package formats

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

func (s *Formats) Create(ctx context.Context, params *CreateParams) (id string, err error) {
	authMetadata, err := s.auth.GetAuthMetadata(params.AccessToken)
	if err != nil {
		return "", err
	}
	if err := s.auth.CheckPermissions(&authMetadata, auth.MoviesFormatsCreate); err != nil {
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

	format, err := s.db.CreateMoviesFormat(ctx, &database.CreateMoviesFormatParams{
		ID:          id,
		DisplayName: params.DisplayName,
		About:       params.About,
	})
	if err != nil {
		return "", err
	}

	formats := make([]database.MoviesFormat, len(s.cache.formats))
	copy(formats, s.cache.formats)

	formats = append(formats, format)
	s.cache.refresh(formats)
	return id, nil
}

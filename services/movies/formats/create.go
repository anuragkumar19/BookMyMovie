package formats

import (
	"context"
	"strings"

	"bookmymovie.app/bookmymovie/database"
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

func (params *CreateParams) Validate() error {
	return validation.ValidateStruct(params,
		validation.Field(&params.DisplayName, validation.Required),
		validation.Field(&params.About, validation.Required),
	)
}

func (s *Formats) Create(ctx context.Context, authMeta *auth.Metadata, params *CreateParams) (database.MoviesFormat, error) {
	if err := authMeta.Valid(); err != nil {
		return database.MoviesFormat{}, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesFormatsCreate); err != nil {
		return database.MoviesFormat{}, err
	}

	slg := slug.Make(params.DisplayName)

	format, err := s.db.CreateMoviesFormat(ctx, &database.CreateMoviesFormatParams{
		Slug:        slg,
		DisplayName: params.DisplayName,
		About:       params.About,
	})
	if err != nil {
		return database.MoviesFormat{}, err
	}

	formats := make([]database.MoviesFormat, len(s.cache.formats))
	copy(formats, s.cache.formats)

	formats = append(formats, format)
	s.cache.refresh(formats)
	return format, nil
}

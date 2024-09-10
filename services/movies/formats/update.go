package formats

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
	About       *string
}

func (params *UpdateParams) Transform() *UpdateParams {
	if params.DisplayName != nil {
		dn := strings.TrimSpace(*params.DisplayName)
		params.DisplayName = &dn
	}
	if params.About != nil {
		abt := strings.TrimSpace(*params.About)
		params.About = &abt
	}
	return params
}

func (params *UpdateParams) Validate() error {
	return validation.ValidateStruct(params,
		validation.Field(&params.DisplayName, validation.Length(1, 0)),
		validation.Field(&params.About, validation.Length(1, 0)),
	)
}

func (s *Formats) Update(ctx context.Context, authMeta *auth.Metadata, params *UpdateParams) (database.MoviesFormat, error) {
	if err := authMeta.Valid(); err != nil {
		return database.MoviesFormat{}, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesFormatsUpdate); err != nil {
		return database.MoviesFormat{}, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return database.MoviesFormat{}, err
		}
		return database.MoviesFormat{}, serviceserrors.New(serviceserrors.ErrorTypeInvalidArgument, err.Error())
	}

	format, err := s.GetByID(ctx, params.ID)
	if err != nil {
		return database.MoviesFormat{}, err
	}

	if params.DisplayName != nil {
		format.DisplayName = *params.DisplayName
	}
	if params.About != nil {
		format.About = *params.About
	}
	format.Slug = slug.Make(format.DisplayName)
	if err := s.db.UpdateMoviesFormat(ctx, &database.UpdateMoviesFormatParams{
		Slug:        format.Slug,
		DisplayName: format.DisplayName,
		About:       format.About,
		ID:          format.ID,
		Version:     format.Version,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return database.MoviesFormat{}, serviceserrors.New(serviceserrors.ErrorConflict, "")
		}
		return database.MoviesFormat{}, err
	}

	return format, nil
}

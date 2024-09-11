package genres

import (
	"context"
	"errors"
	"strings"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	"bookmymovie.app/bookmymovie/services/auth"
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

func (params UpdateParams) Validate() error {
	return validation.ValidateStruct(&params,
		validation.Field(&params.DisplayName, validation.Length(1, 0)),
		validation.Field(&params.About, validation.Length(1, 0)),
	)
}

func (s *Genres) Update(ctx context.Context, authMeta *auth.Metadata, params *UpdateParams) (database.MoviesGenre, error) {
	if err := authMeta.Valid(); err != nil {
		return database.MoviesGenre{}, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesGenresUpdate); err != nil {
		return database.MoviesGenre{}, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return database.MoviesGenre{}, err
		}
		return database.MoviesGenre{}, services.NewError(services.ErrorTypeInvalidArgument, err.Error())
	}

	genre, err := s.GetByID(ctx, params.ID)
	if err != nil {
		return database.MoviesGenre{}, err
	}

	if params.DisplayName != nil {
		genre.DisplayName = *params.DisplayName
	}
	if params.About != nil {
		genre.About = *params.About
	}
	genre.Slug = slug.Make(genre.DisplayName)
	if err := s.db.UpdateMoviesGenre(ctx, &database.UpdateMoviesGenreParams{
		Slug:        genre.Slug,
		DisplayName: genre.DisplayName,
		About:       genre.About,
		ID:          genre.ID,
		Version:     genre.Version,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return database.MoviesGenre{}, services.NewError(services.ErrorTypeConflict, "")
		}
		return database.MoviesGenre{}, err
	}

	return genre, nil
}

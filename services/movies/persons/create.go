package persons

import (
	"context"
	"strings"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	"bookmymovie.app/bookmymovie/services/auth"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateParams struct {
	Name              string
	Nicknames         []string
	Occupations       []string
	About             string
	ProfilePictureKey string
	DOB               *time.Time
}

func (params *CreateParams) Transform() *CreateParams {
	params.Name = strings.TrimSpace(params.Name)

	params.Nicknames = services.RemoveDuplicate(params.Nicknames)
	params.Nicknames = services.TrimSlices(params.Nicknames)

	params.Occupations = services.RemoveDuplicate(params.Occupations)
	params.Occupations = services.TrimSlices(params.Occupations)

	if params.About == "" {
		params.About = "No information available"
	}

	return params
}

func (params CreateParams) Validate() error {
	return validation.ValidateStruct(
		&params,
		validation.Field(&params.Name, validation.Required),
		validation.Field(&params.About, validation.Required),
		validation.Field(&params.Nicknames, validation.NotNil),
		validation.Field(&params.Occupations, validation.NotNil),
		validation.Field(&params.ProfilePictureKey),
		validation.Field(&params.DOB, validation.Max(time.Now()).Error("dob cannot be in future")),
	)
}

func (s *Persons) Create(ctx context.Context, authMeta *auth.Metadata, params *CreateParams) (id int64, err error) {
	if err := authMeta.Valid(); err != nil {
		return 0, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesPersonsCreate); err != nil {
		return 0, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return 0, err
		}
		return 0, services.NewError(services.ErrorTypeInvalidArgument, err.Error())
	}

	exist, err := s.storage.Exist(ctx, params.ProfilePictureKey)
	if err != nil {
		return 0, err
	}
	if !exist {
		return 0, services.NewError(services.ErrorTypeInvalidArgument, "profile picture selected doesn't exit")
	}

	slg := slug.Make(params.Name)

	return s.db.CreateMoviesPerson(ctx, &database.CreateMoviesPersonParams{
		Name:           params.Name,
		Slug:           slg,
		Nicknames:      params.Nicknames,
		ProfilePicture: params.ProfilePictureKey,
		Occupations:    params.Occupations,
		Dob: pgtype.Date{
			Valid: params.DOB != nil,
			Time:  *params.DOB,
		},
		About: params.About,
	})
}

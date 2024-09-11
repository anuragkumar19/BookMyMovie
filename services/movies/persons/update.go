package persons

import (
	"context"
	"errors"
	"strings"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	"bookmymovie.app/bookmymovie/services/auth"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateParams struct {
	ID                int64
	Name              *string
	ProfilePictureKey *string
	Dob               *time.Time
	// TODO: look for possibility of doing all updates in same function
	AppendOccupations []string
	RemoveOccupations []string
	AppendNickname    []string
	RemoveNickname    []string
}

func (params *UpdateParams) Transform() *UpdateParams {
	if params.Name != nil {
		nm := strings.TrimSpace(*params.Name)
		params.Name = &nm
	}

	return params
}

func (params UpdateParams) Validate() error {
	return validation.ValidateStruct(
		&params,
		validation.Field(&params.Name, validation.Min(1)),
		validation.Field(&params.ProfilePictureKey),
		validation.Field(&params.Dob, validation.Max(time.Now()).Error("dob cannot be in future")),
	)
}

func (s *Persons) Update(ctx context.Context, authMeta *auth.Metadata, params *UpdateParams) (database.GetMoviesPersonRow, error) {
	if err := authMeta.Valid(); err != nil {
		return database.GetMoviesPersonRow{}, err
	}
	if err := s.auth.CheckPermissions(authMeta, auth.MoviesPersonsCreate); err != nil {
		return database.GetMoviesPersonRow{}, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return database.GetMoviesPersonRow{}, err
		}
		return database.GetMoviesPersonRow{}, services.NewError(services.ErrorTypeInvalidArgument, err.Error())
	}

	person, err := s.db.GetMoviesPerson(ctx, params.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return database.GetMoviesPersonRow{}, services.NewError(services.ErrorTypeNotFound, "movies person not found")
		}
	}

	if params.Name != nil {
		person.Name = *params.Name
		person.Slug = slug.Make(*params.Name)
	}
	if params.Dob != nil {
		person.Dob = pgtype.Date{
			Valid: !params.Dob.IsZero(),
			Time:  *params.Dob,
		}
	}
	if params.ProfilePictureKey != nil {
		exist, err := s.storage.Exist(ctx, *params.ProfilePictureKey)
		if err != nil {
			return database.GetMoviesPersonRow{}, err
		}
		if !exist {
			return database.GetMoviesPersonRow{}, services.NewError(services.ErrorTypeInvalidArgument, "profile picture selected doesn't exit")
		}

		person.ProfilePicture = *params.ProfilePictureKey
	}

	if err := s.db.UpdateMoviesPerson(ctx, &database.UpdateMoviesPersonParams{
		Name:           person.Name,
		Slug:           person.Slug,
		Nicknames:      person.Nicknames,
		ProfilePicture: person.ProfilePicture,
		Occupations:    person.Occupations,
		Dob:            person.Dob,
		About:          person.About,
		ID:             person.ID,
		Version:        person.Version,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return database.GetMoviesPersonRow{}, services.NewError(services.ErrorTypeConflict, "")
		}
		return database.GetMoviesPersonRow{}, err
	}
	return person, nil
}

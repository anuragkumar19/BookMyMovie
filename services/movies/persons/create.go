package persons

import (
	"context"
	"errors"
	"strings"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateParams struct {
	AccessToken string

	Name              string
	Nicknames         []string
	Occupations       []string
	About             string
	ProfilePictureKey string
	DOB               *time.Time
	ImdbID            string
}

func (params *CreateParams) Transform() *CreateParams {
	params.Name = strings.TrimSpace(params.Name)
	nicknames := []string{}
	for _, n := range params.Nicknames {
		nicknames = append(nicknames, strings.TrimSpace(n))
	}
	params.Nicknames = nicknames
	occupations := []string{}
	for _, n := range params.Occupations {
		occupations = append(occupations, strings.TrimSpace(n))
	}
	params.Occupations = occupations
	if params.About == "" {
		params.About = "No information available"
	}
	return params
}

func (params *CreateParams) Validate() error {
	return validation.ValidateStruct(
		params,
		validation.Field(&params.Name, validation.Required),
		validation.Field(&params.About, validation.Required),
		validation.Field(&params.Nicknames),
		validation.Field(&params.Occupations),
		validation.Field(&params.ProfilePictureKey),
		validation.Field(&params.DOB, validation.Max(time.Now()).Error("dob cannot be in future")),
		validation.Field(&params.ImdbID, validation.Required),
	)
}

func (s *Persons) Create(ctx context.Context, params *CreateParams) (id int64, err error) {
	authMetadata, err := s.auth.GetAuthMetadata(params.AccessToken)
	if err != nil {
		return 0, err
	}
	if err := s.auth.CheckPermissions(&authMetadata, auth.MoviesPersonsCreate); err != nil {
		return 0, err
	}

	exist, err := s.storage.Exist(ctx, params.ProfilePictureKey)
	if err != nil {
		return 0, err
	}
	if !exist {
		return 0, errors.New("profile picture selected doesn't exit") // TODO: better error
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
		About:            params.About,
		ImdbID:           params.ImdbID,
		ImdbLastSyncedAt: pgtype.Timestamptz{Valid: false},
	})
}

package users

import (
	"context"
	"errors"
	"strings"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	"bookmymovie.app/bookmymovie/services/auth"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UpdateParams struct {
	Name *string
	Dob  *time.Time
}

func (params *UpdateParams) Transform() *UpdateParams {
	if params.Name != nil {
		*params.Name = strings.TrimSpace(*params.Name)
	}
	return params
}

func (params UpdateParams) Validate() error {
	return validation.ValidateStruct(&params,
		validation.Field(&params.Name, validation.Min(1)),
		validation.Field(&params.Dob, validation.Max(time.Now().Add(-5*365*24*time.Hour)).Error("age must be at least 5 years")),
	)
}

func (s *Users) Update(ctx context.Context, authMeta *auth.Metadata, params *UpdateParams) (database.FindUserByIdRow, error) {
	if err := authMeta.Valid(); err != nil {
		return database.FindUserByIdRow{}, err
	}

	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return database.FindUserByIdRow{}, err
		}
		return database.FindUserByIdRow{}, services.NewError(services.ErrorTypeInvalidArgument, err.Error())
	}
	user, err := s.db.FindUserById(ctx, authMeta.UserID())
	if err != nil {
		return database.FindUserByIdRow{}, err
	}

	if params.Name != nil {
		user.Name = *params.Name
	}
	if params.Dob != nil {
		user.Dob = pgtype.Date{
			Valid: params.Dob.IsZero(),
			Time:  *params.Dob,
		}
	}

	if err := s.db.UpdateUserProfile(ctx, &database.UpdateUserProfileParams{
		Name:    user.Name,
		Dob:     user.Dob,
		ID:      user.ID,
		Version: user.Version,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return database.FindUserByIdRow{}, services.NewError(services.ErrorTypeConflict, "")
		}
		return database.FindUserByIdRow{}, err
	}

	return user, nil
}

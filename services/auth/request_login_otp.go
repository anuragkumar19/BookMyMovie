package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"bookmymovie.app/bookmymovie/database"
	services_errors "bookmymovie.app/bookmymovie/services/serviceserrors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type RequestLoginOTPParams struct {
	Email string
}

func (data *RequestLoginOTPParams) Transform() *RequestLoginOTPParams {
	data.Email = strings.TrimSpace(data.Email)
	return data
}

func (data *RequestLoginOTPParams) Validate() error {
	return validation.ValidateStruct(
		data,
		validation.Field(&data.Email, validation.Required, is.Email),
	)
}

func (s *Auth) RequestLoginOTP(ctx context.Context, params *RequestLoginOTPParams) (loginToken string, err error) {
	if err := params.Transform().Validate(); err != nil {
		return "", services_errors.ValidationError(err.(validation.Errors)) //nolint:errorlint
	}

	var isNew bool
	user, err := s.db.FindUserByEmail(ctx, params.Email)

	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return "", err
		}

		id, err := s.db.CreateRegularUser(ctx, params.Email)
		if err != nil {
			return "", err
		}
		user.ID = id
		user.Email = params.Email
		user.Version = 1
		isNew = true
	}

	now := time.Now()
	nowMinusWindow := now.Add(-s.config.LoginOTPSendingRateTimeWindow)
	if user.LastLoginTokenSentAt.Time.Before(nowMinusWindow) {
		user.TotalLoginTokensSent = 0
	}
	if user.TotalLoginTokensSent >= int32(s.config.LoginOTPSendingRate) {
		return "", services_errors.NewRateLimitError(user.LastLoginTokenSentAt.Time.Add(s.config.LoginOTPSendingRateTimeWindow).Sub(now), int(user.TotalLoginTokensSent), user.LastLoginTokenSentAt.Time)
	}

	token, err := s.generateRandomToken()
	if err != nil {
		return "", err
	}
	otp, otpHash, err := s.generateOTP()
	if err != nil {
		return "", err
	}

	expiry := now.Add(s.config.LoginOTPLifetime)
	if err := s.db.CreateLoginToken(ctx, &database.CreateLoginTokenParams{
		Token:     token,
		Otp:       otpHash,
		CreatedAt: pgtype.Timestamptz{Time: now, Valid: true},
		ExpireAt:  pgtype.Timestamptz{Time: expiry, Valid: true},
		UserID:    user.ID,
	}); err != nil {
		return "", err
	}

	if err := s.db.UpdateUserLoginFields(ctx, &database.UpdateUserLoginFieldsParams{
		LastLoginTokenSentAt: pgtype.Timestamptz{Time: now, Valid: true},
		TotalLoginTokensSent: user.TotalLoginTokensSent + 1,
		ID:                   user.ID,
		Version:              user.Version,
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", services_errors.ErrUpdateConflict
		}
		return "", err
	}

	// TODO: send token+otp link and just otp
	s.logger.Info().Str("email", user.Email).Str("otp", otp).Str("link", s.generateLoginLink(token, otp)).Time("expire_at", expiry).Bool("is_new", isNew).Msg("mail sent")

	return token, nil
}

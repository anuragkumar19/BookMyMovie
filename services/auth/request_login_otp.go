package auth

import (
	"context"
	"errors"
	"strings"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
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
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return "", err
		}
		return "", serviceserrors.New(serviceserrors.ErrorTypeInvalidArgument, err.Error())
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
		return "", serviceserrors.New(serviceserrors.ErrorResourceExhausted, serviceserrors.NewRateLimitMessage(time.Until(user.LastLoginTokenSentAt.Time.Add(s.config.LoginOTPSendingRateTimeWindow)), ""))
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
			return "", serviceserrors.New(serviceserrors.ErrorConflict, "")
		}
		return "", err
	}

	link := s.generateLoginLink(token, otp)
	msg, err := mailer.NewLoginMessage(ctx, &mailer.LoginMessageParams{
		OTP:        otp,
		Link:       link,
		IsNew:      isNew,
		Email:      user.Email,
		ExpiryTime: s.config.LoginOTPLifetime,
	})
	if err != nil {
		return "", err
	}
	s.mailer.SendMessage(&msg)

	return token, nil
}

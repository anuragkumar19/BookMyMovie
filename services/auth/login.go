package auth

import (
	"context"
	"net/url"
	"strings"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type LoginParams struct {
	Token string `json:"token"`
	OTP   string `json:"otp"`
}

func (data *LoginParams) Transform() *LoginParams {
	data.OTP = strings.TrimSpace(data.OTP)
	data.Token = strings.TrimSpace(data.Token)
	return data
}

func (data *LoginParams) Validate() error {
	return validation.ValidateStruct(
		data,
		validation.Field(&data.Token, validation.Required),
		validation.Field(&data.OTP, validation.Required),
	)
}

type AuthTokens struct {
	RefreshToken      string
	AccessToken       string
	AccessTokenExpiry time.Time
}

func (s *Auth) Login(ctx context.Context, params *LoginParams) (AuthTokens, error) {
	if err := params.Transform().Validate(); err != nil {
		return AuthTokens{}, err
	}

	token, err := s.db.FindLoginToken(ctx, params.Token)
	if err != nil {
		if err == pgx.ErrNoRows {
			return AuthTokens{}, errors.ErrOTPExpired
		}
		return AuthTokens{}, err
	}

	now := time.Now()

	if token.ExpireAt.Time.Before(now) {
		return AuthTokens{}, errors.ErrOTPExpired
	}
	if token.TotalAttempts >= int32(s.config.MaxOTPIncorrectAttempts) {
		if err := s.db.DeleteLoginToken(ctx, token.Token); err != nil {
			if err == pgx.ErrNoRows {
				return AuthTokens{}, errors.ErrOTPExpired
			}
			return AuthTokens{}, err
		}
		return AuthTokens{}, errors.ErrOTPExpired
	}

	if !s.matchOTP(params.OTP, token.Otp) {
		if err := s.db.AttemptLoginToken(ctx, &database.AttemptLoginTokenParams{
			LastAttemptAt: pgtype.Timestamptz{Time: now, Valid: true},
			TotalAttempts: token.TotalAttempts + 1,
			Token:         token.Token,
			Version:       token.Version,
		}); err != nil {
			if err == pgx.ErrNoRows {
				return AuthTokens{}, errors.ErrUpdateConflict
			}
			return AuthTokens{}, err
		}
		return AuthTokens{}, errors.ErrOTPMismatch
	}

	if err := s.db.DeleteLoginToken(ctx, token.Token); err != nil {
		if err == pgx.ErrNoRows {
			return AuthTokens{}, errors.ErrOTPExpired
		}
		return AuthTokens{}, err
	}

	// TODO: generate PASETO tokens
	return AuthTokens{}, nil
}

func (s *Auth) generateLoginLink(token string, otp string) string {
	return s.config.Host + "/auth/login?token=" + url.QueryEscape(token) + "&otp=" + url.QueryEscape(otp)
}

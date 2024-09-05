package auth

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type LoginParams struct {
	Token     string
	OTP       string
	UserAgent string
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
		validation.Field(&data.UserAgent, validation.Required),
	)
}

type Tokens struct {
	RefreshToken      string
	AccessToken       string
	AccessTokenExpiry time.Time
}

func (s *Auth) Login(ctx context.Context, params *LoginParams) (Tokens, error) {
	if err := params.Transform().Validate(); err != nil {
		return Tokens{}, serviceserrors.ValidationError(err.(validation.Errors)) //nolint:errorlint
	}

	token, err := s.db.FindLoginToken(ctx, params.Token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Tokens{}, serviceserrors.ErrOTPExpired
		}
		return Tokens{}, err
	}

	now := time.Now()

	if token.ExpireAt.Time.Before(now) {
		return Tokens{}, serviceserrors.ErrOTPExpired
	}
	if token.TotalAttempts >= int32(s.config.MaxOTPIncorrectAttempts) {
		if err := s.db.DeleteLoginToken(ctx, token.Token); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return Tokens{}, serviceserrors.ErrOTPExpired
			}
			return Tokens{}, err
		}
		return Tokens{}, serviceserrors.ErrOTPExpired
	}

	if !s.matchOTP(params.OTP, token.Otp) {
		if err := s.db.AttemptLoginToken(ctx, &database.AttemptLoginTokenParams{
			LastAttemptAt: pgtype.Timestamptz{Time: now, Valid: true},
			TotalAttempts: token.TotalAttempts + 1,
			Token:         token.Token,
			Version:       token.Version,
		}); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return Tokens{}, serviceserrors.ErrUpdateConflict
			}
			return Tokens{}, err
		}
		return Tokens{}, serviceserrors.ErrOTPMismatch
	}

	if err := s.db.DeleteLoginToken(ctx, token.Token); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Tokens{}, serviceserrors.ErrOTPExpired
		}
		return Tokens{}, err
	}

	refreshTokenStr, err := s.generateRandomToken()
	if err != nil {
		return Tokens{}, err
	}
	refreshToken, err := s.db.CreateRefreshToken(ctx, &database.CreateRefreshTokenParams{
		Token:     refreshTokenStr,
		CreatedAt: pgtype.Timestamptz{Time: now, Valid: true},
		UserID:    token.UserID,
		UserRole:  token.UserRole,
		ExpireAt:  pgtype.Timestamptz{Time: now.Add(s.config.RefreshTokenLifetime), Valid: true},
		UserAgent: params.UserAgent,
	})
	if err != nil {
		return Tokens{}, err
	}

	accessToken, accessTokenExp, err := s.generateAccessToken(&refreshToken)
	if err != nil {
		return Tokens{}, err
	}
	return Tokens{
		RefreshToken:      refreshTokenStr,
		AccessToken:       accessToken,
		AccessTokenExpiry: accessTokenExp,
	}, nil
}

func (s *Auth) generateLoginLink(token string, otp string) string {
	return s.config.AppPublicHost + "/auth/login?token=" + url.QueryEscape(token) + "&otp=" + url.QueryEscape(otp)
}

package auth

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var errOTPExpired = errors.New("otp expired")
var errOTPMismatched = errors.New("otp mismatched")

type LoginParams struct {
	Token     string
	OTP       string
	UserAgent string
}

func (params *LoginParams) Transform() *LoginParams {
	params.OTP = strings.TrimSpace(params.OTP)
	params.Token = strings.TrimSpace(params.Token)
	return params
}

func (params LoginParams) Validate() error {
	return validation.ValidateStruct(
		&params,
		validation.Field(&params.Token, validation.Required),
		validation.Field(&params.OTP, validation.Required),
		validation.Field(&params.UserAgent, validation.Required),
	)
}

type Tokens struct {
	RefreshToken      string
	AccessToken       string
	AccessTokenExpiry time.Time
}

func (s *Auth) Login(ctx context.Context, params *LoginParams) (Tokens, error) {
	if err := params.Transform().Validate(); err != nil {
		if _, ok := err.(validation.InternalError); ok { //nolint:errorlint
			return Tokens{}, err
		}
		return Tokens{}, services.NewError(services.ErrorTypeInvalidArgument, err.Error())
	}

	token, err := s.db.FindLoginToken(ctx, params.Token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Tokens{}, services.NewError(services.ErrorTypeInvalidArgument, errOTPExpired.Error())
		}
		return Tokens{}, err
	}

	now := time.Now()

	if token.ExpireAt.Time.Before(now) {
		return Tokens{}, services.NewError(services.ErrorTypeInvalidArgument, errOTPExpired.Error())
	}
	if token.TotalAttempts >= int32(s.config.MaxOTPIncorrectAttempts) {
		if err := s.db.DeleteLoginToken(ctx, token.Token); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return Tokens{}, services.NewError(services.ErrorTypeInvalidArgument, errOTPExpired.Error())
			}
			return Tokens{}, err
		}
		return Tokens{}, services.NewError(services.ErrorTypeInvalidArgument, errOTPExpired.Error())
	}

	if !s.matchOTP(params.OTP, token.Otp) {
		if err := s.db.AttemptLoginToken(ctx, &database.AttemptLoginTokenParams{
			LastAttemptAt: pgtype.Timestamptz{Time: now, Valid: true},
			TotalAttempts: token.TotalAttempts + 1,
			Token:         token.Token,
			Version:       token.Version,
		}); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return Tokens{}, services.NewError(services.ErrorTypeConflict, "")
			}
			return Tokens{}, err
		}
		return Tokens{}, services.NewError(services.ErrorTypeInvalidArgument, errOTPMismatched.Error())
	}

	if err := s.db.DeleteLoginToken(ctx, token.Token); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Tokens{}, services.NewError(services.ErrorTypeInvalidArgument, errOTPExpired.Error())
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

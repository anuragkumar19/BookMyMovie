package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"bookmymovie.app/bookmymovie/database"
	services_errors "bookmymovie.app/bookmymovie/services/errors"
	"github.com/anuragkumar19/binding"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type LoginParams struct {
	Token     string `json:"token"`
	OTP       string `json:"otp"`
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

type AuthTokens struct {
	RefreshToken      string
	AccessToken       string
	AccessTokenExpiry time.Time
}

func (s *Auth) Login(ctx context.Context, params *LoginParams) (AuthTokens, error) {
	if err := params.Transform().Validate(); err != nil {
		return AuthTokens{}, services_errors.ValidationError(err.(validation.Errors))
	}

	token, err := s.db.FindLoginToken(ctx, params.Token)
	if err != nil {
		if err == pgx.ErrNoRows {
			return AuthTokens{}, services_errors.ErrOTPExpired
		}
		return AuthTokens{}, err
	}

	now := time.Now()

	if token.ExpireAt.Time.Before(now) {
		return AuthTokens{}, services_errors.ErrOTPExpired
	}
	if token.TotalAttempts >= int32(s.config.MaxOTPIncorrectAttempts) {
		if err := s.db.DeleteLoginToken(ctx, token.Token); err != nil {
			if err == pgx.ErrNoRows {
				return AuthTokens{}, services_errors.ErrOTPExpired
			}
			return AuthTokens{}, err
		}
		return AuthTokens{}, services_errors.ErrOTPExpired
	}

	if !s.matchOTP(params.OTP, token.Otp) {
		if err := s.db.AttemptLoginToken(ctx, &database.AttemptLoginTokenParams{
			LastAttemptAt: pgtype.Timestamptz{Time: now, Valid: true},
			TotalAttempts: token.TotalAttempts + 1,
			Token:         token.Token,
			Version:       token.Version,
		}); err != nil {
			if err == pgx.ErrNoRows {
				return AuthTokens{}, services_errors.ErrUpdateConflict
			}
			return AuthTokens{}, err
		}
		return AuthTokens{}, services_errors.ErrOTPMismatch
	}

	if err := s.db.DeleteLoginToken(ctx, token.Token); err != nil {
		if err == pgx.ErrNoRows {
			return AuthTokens{}, services_errors.ErrOTPExpired
		}
		return AuthTokens{}, err
	}

	refreshTokenStr, err := s.generateRandomToken()
	if err != nil {
		return AuthTokens{}, err
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
		return AuthTokens{}, err
	}

	accessToken, accessTokenExp, err := s.generateAccessToken(&refreshToken)
	if err != nil {
		return AuthTokens{}, err
	}
	return AuthTokens{
		RefreshToken:      refreshTokenStr,
		AccessToken:       accessToken,
		AccessTokenExpiry: accessTokenExp,
	}, nil
}

func (s *Auth) loginHandler(w http.ResponseWriter, r *http.Request) {
	var params LoginParams
	if err := binding.Bind(r, &params); err != nil {
		services_errors.HTTPErrorHandler(err, w, r)
		return
	}
	params.UserAgent = r.UserAgent()

	tks, err := s.Login(r.Context(), &params)
	if err != nil {
		services_errors.HTTPErrorHandler(err, w, r)
		return
	}
	w.Write([]byte(fmt.Sprintf("%#v", tks)))
}

func (s *Auth) generateLoginLink(token string, otp string) string {
	return s.config.Host + "/auth/login?token=" + url.QueryEscape(token) + "&otp=" + url.QueryEscape(otp)
}

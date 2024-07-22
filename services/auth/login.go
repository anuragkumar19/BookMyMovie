package auth

import (
	"context"
	"errors"
	"time"

	"bookmymovie.app/bookmymovie/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/oklog/ulid/v2"
)

const loginRateLimitWindow = time.Hour
const loginRateLimit = 10
const loginTokenExpiry = 10 * time.Minute

const allowedIncorrectOTPAttempts = 10

type RequestLoginOTPParams struct {
	Email string `json:"email"`
}

func (data *RequestLoginOTPParams) Validate() error {
	return validation.ValidateStruct(
		data,
		validation.Field(&data.Email, validation.Required, is.Email),
	)
}

func (s *Auth) RequestLoginOTP(ctx context.Context, params *RequestLoginOTPParams) (loginToken string, err error) {
	if err := params.Validate(); err != nil {
		return "", err
	}

	var isNew bool
	user, err := s.db.FindUserByEmail(ctx, params.Email)

	if err != nil {
		if err == pgx.ErrNoRows {
			userId := ulid.Make().Bytes()
			err := s.db.CreateRegularUser(ctx, &database.CreateRegularUserParams{
				Email: loginToken,
				ID:    userId,
			})
			if err != nil {
				return "", err
			}
			user.ID = userId
			user.Email = params.Email
			user.Version = 1
			isNew = true
		} else {
			return "", err
		}
	}

	now := time.Now()
	nowMinusWindow := now.Add(-loginRateLimitWindow)
	if user.LastLoginTokenSentAt.Time.Before(nowMinusWindow) {
		user.TotalLoginTokensSent = 0
	}
	if user.TotalLoginTokensSent >= loginRateLimit {
		//TODO: proper error with services/errors package
		return "", errors.New("rate limited")
	}

	token, err := generateRandomToken()
	if err != nil {
		return "", err
	}
	otp, otpHash, err := generateOTP()
	if err != nil {
		return "", err
	}

	expiry := now.Add(loginTokenExpiry)
	if err := s.db.CreateLoginToken(ctx, &database.CreateLoginTokenParams{
		Token:     token,
		Otp:       otpHash,
		CreatedAt: pgtype.Timestamptz{Time: now, Valid: true},
		ExpireAt:  pgtype.Timestamptz{Time: expiry, Valid: true},
		UserID:    user.ID,
	}); err != nil {
		return "", err
	}

	//TODO: send token+otp link and just otp
	s.logger.Info().Str("email", user.Email).Str("otp", otp).Str("token", token).Time("expire_at", expiry).Bool("is_new", isNew).Msg("mail sent")

	if err := s.db.UpdateUserLoginFields(ctx, &database.UpdateUserLoginFieldsParams{
		LastLoginTokenSentAt: pgtype.Timestamptz{Time: now, Valid: true},
		TotalLoginTokensSent: user.TotalLoginTokensSent + 1,
		ID:                   user.ID,
		Version:              user.Version,
	}); err != nil {
		if err == pgx.ErrNoRows {
			//TODO: proper errors
			return "", errors.New("update conflict, please try again")
		} else {
			return "", err
		}
	}

	return token, nil
}

type LoginParams struct {
	Token string `json:"token"`
	OTP   string `json:"otp"`
}

func (data *LoginParams) Validate() error {
	return validation.ValidateStruct(
		data,
		validation.Field(&data.Token, validation.Required),
		validation.Field(&data.OTP, validation.Required, validation.Length(otpLength, otpLength)),
	)
}

type AuthTokens struct {
	RefreshToken      string
	AccessToken       string
	AccessTokenExpiry time.Time
}

func (s *Auth) Login(ctx context.Context, params *LoginParams) (AuthTokens, error) {
	if err := params.Validate(); err != nil {
		return AuthTokens{}, err
	}

	token, err := s.db.FindLoginToken(ctx, params.Token)
	if err != nil {
		if err == pgx.ErrNoRows {
			// TODO: proper error
			return AuthTokens{}, errors.New("otp expired")
		}
		return AuthTokens{}, err
	}

	now := time.Now()

	if token.ExpireAt.Time.Before(now) {
		// TODO: proper error
		return AuthTokens{}, errors.New("otp expired")
	}
	if token.TotalAttempts >= allowedIncorrectOTPAttempts {
		// TODO: proper error
		return AuthTokens{}, errors.New("too many attempts")
	}

	if !matchOTP(params.OTP, token.Otp) {
		if err := s.db.AttemptLoginToken(ctx, &database.AttemptLoginTokenParams{
			LastAttemptAt: pgtype.Timestamptz{Time: now, Valid: true},
			TotalAttempts: token.TotalAttempts + 1,
			Token:         token.Token,
			Version:       token.Version,
		}); err != nil {
			if err == pgx.ErrNoRows {
				//TODO: proper errors
				return AuthTokens{}, errors.New("update conflict, please try again")
			}
			return AuthTokens{}, err
		}
		//TODO: proper errors
		return AuthTokens{}, errors.New("otp mismatch")
	}

	if err := s.db.DeleteLoginToken(ctx, token.Token); err != nil {
		if err == pgx.ErrNoRows {
			return AuthTokens{}, errors.New("otp expired")
		}
		return AuthTokens{}, err
	}

	// TODO: generate PASETO tokens
	return AuthTokens{}, nil
}

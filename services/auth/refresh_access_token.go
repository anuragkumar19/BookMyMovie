package auth

import (
	"context"
	"errors"
	"time"

	"bookmymovie.app/bookmymovie/services/serviceserrors"
	"github.com/jackc/pgx/v5"
)

type AccessToken struct {
	AccessToken       string
	AccessTokenExpiry time.Time
}

func (s *Auth) RefreshAccessToken(ctx context.Context, token string) (AccessToken, error) {
	if token == "" {
		return AccessToken{}, serviceserrors.UnauthorizedError(ErrTokenInvalid)
	}

	rt, err := s.db.FindRefreshToken(ctx, token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return AccessToken{}, serviceserrors.UnauthorizedError(ErrTokenInvalid)
		}
	}
	if time.Now().After(rt.ExpireAt.Time) {
		return AccessToken{}, serviceserrors.UnauthorizedError(ErrTokenInvalid)
	}

	accessToken, accessTokenExp, err := s.generateAccessToken(&rt)
	if err != nil {
		return AccessToken{}, err
	}
	return AccessToken{
		AccessToken:       accessToken,
		AccessTokenExpiry: accessTokenExp,
	}, nil
}

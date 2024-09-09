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
		return AccessToken{}, serviceserrors.New(serviceserrors.ErrorUnauthenticated, errTokenInvalid.Error())
	}

	rt, err := s.db.FindRefreshToken(ctx, token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return AccessToken{}, serviceserrors.New(serviceserrors.ErrorUnauthenticated, errTokenInvalid.Error())
		}
	}
	if time.Now().After(rt.ExpireAt.Time) {
		return AccessToken{}, serviceserrors.New(serviceserrors.ErrorUnauthenticated, errTokenInvalid.Error())
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

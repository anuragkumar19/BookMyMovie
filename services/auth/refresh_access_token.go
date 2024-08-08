package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	services_errors "bookmymovie.app/bookmymovie/services/errors"
	"github.com/jackc/pgx/v5"
)

type AccessToken struct {
	AccessToken       string
	AccessTokenExpiry time.Time
}

func (s *Auth) RefreshAccessToken(ctx context.Context, token string) (AccessToken, error) {
	if token == "" {
		return AccessToken{}, services_errors.UnauthorizedError(ErrTokenInvalid)
	}

	rt, err := s.db.FindRefreshToken(ctx, token)
	if err != nil {
		if err == pgx.ErrNoRows {
			return AccessToken{}, services_errors.UnauthorizedError(ErrTokenInvalid)
		}
	}
	if time.Now().After(rt.ExpireAt.Time) {
		return AccessToken{}, services_errors.UnauthorizedError(ErrTokenInvalid)
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

func (s *Auth) refreshAccessTokenHandler(w http.ResponseWriter, r *http.Request) {

	token, err := s.RefreshAccessToken(r.Context(), r.Header.Get("Authorization"))
	if err != nil {
		services_errors.HTTPErrorHandler(err, w, r)
		return
	}

	w.Write([]byte(fmt.Sprintf("%#v", token)))
}

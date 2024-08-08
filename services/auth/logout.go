package auth

import (
	"context"
	"net/http"
	"time"

	services_errors "bookmymovie.app/bookmymovie/services/errors"
	"github.com/jackc/pgx/v5"
)

func (s *Auth) Logout(ctx context.Context, accessToken string) error {
	auth, err := s.GetAuthMetadata(accessToken)
	if err != nil {
		return err
	}

	if err := s.db.DeleteRefreshToken(ctx, auth.RefreshTokenID); err != nil {
		if err == pgx.ErrNoRows {
			return services_errors.ErrUpdateConflict
		}
		return err
	}
	s.revokedTokens[auth.RefreshTokenID] = time.Now().Add(s.config.AccessTokenLifetime)
	return nil
}

func (s *Auth) logoutHandler(w http.ResponseWriter, r *http.Request) {

	err := s.Logout(r.Context(), r.Header.Get("Authorization"))
	if err != nil {
		services_errors.HTTPErrorHandler(err, w, r)
		return
	}

	w.WriteHeader(200)
}

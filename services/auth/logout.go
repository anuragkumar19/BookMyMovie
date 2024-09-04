package auth

import (
	"context"
	"errors"
	"time"

	"bookmymovie.app/bookmymovie/services/serviceserrors"
	"github.com/jackc/pgx/v5"
)

func (s *Auth) Logout(ctx context.Context, accessToken string) error {
	auth, err := s.GetAuthMetadata(accessToken)
	if err != nil {
		return err
	}

	if err := s.db.DeleteRefreshToken(ctx, auth.RefreshTokenID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return serviceserrors.ErrUpdateConflict
		}
		return err
	}
	s.revokedTokens[auth.RefreshTokenID] = time.Now().Add(s.config.AccessTokenLifetime)
	return nil
}

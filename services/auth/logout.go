package auth

import (
	"context"
	"errors"
	"time"

	"bookmymovie.app/bookmymovie/services"
	"github.com/jackc/pgx/v5"
)

func (s *Auth) Logout(ctx context.Context, authMeta *Metadata) error {
	if err := authMeta.Valid(); err != nil {
		return err
	}
	if err := s.db.DeleteRefreshToken(ctx, authMeta.RefreshTokenID()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return services.NewError(services.ErrorTypeConflict, "")
		}
		return err
	}
	s.revokedTokens[authMeta.RefreshTokenID()] = time.Now().Add(s.config.AccessTokenLifetime)
	return nil
}

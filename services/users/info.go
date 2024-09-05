package users

import (
	"context"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/auth"
)

func (s *Users) GetLoggedInUser(ctx context.Context, authMeta *auth.Metadata) (database.FindUserByIdRow, error) {
	if err := authMeta.Valid(); err != nil {
		return database.FindUserByIdRow{}, err
	}
	user, err := s.db.FindUserById(ctx, authMeta.UserID())
	if err != nil {
		return database.FindUserByIdRow{}, err
	}
	return user, nil
}

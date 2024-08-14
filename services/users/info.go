package users

import (
	"context"

	"bookmymovie.app/bookmymovie/database"
)

func (s *Users) GetUserInfo(ctx context.Context, accessToken string) (database.FindUserByIdRow, error) {
	authMeta, err := s.auth.GetAuthMetadata(accessToken)
	if err != nil {
		return database.FindUserByIdRow{}, err
	}
	user, err := s.db.FindUserById(ctx, authMeta.UserID)
	if err != nil {
		return database.FindUserByIdRow{}, err
	}
	return user, nil
}

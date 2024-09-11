package api

import (
	"context"

	"bookmymovie.app/bookmymovie"
	usersv1 "bookmymovie.app/bookmymovie/api/gen/users/v1"
	"bookmymovie.app/bookmymovie/api/gen/users/v1/usersv1connect"
	"bookmymovie.app/bookmymovie/services/users"
	"connectrpc.com/connect"
)

type usersService struct {
	usersv1connect.UnimplementedUsersServiceHandler

	app *bookmymovie.Application
}

func (s *usersService) GetLoggedInUser(ctx context.Context, r *connect.Request[usersv1.GetLoggedInUserRequest]) (*connect.Response[usersv1.GetLoggedInUserResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	user, err := s.app.UsersService().GetLoggedInUser(ctx, &authMeta)
	if err != nil {
		return nil, serviceErrorHandler(err)
	}
	res := connect.NewResponse(&usersv1.GetLoggedInUserResponse{
		User: mapUser(&user),
	})
	return res, nil
}

func (s *usersService) UpdateUser(ctx context.Context, r *connect.Request[usersv1.UpdateUserRequest]) (*connect.Response[usersv1.UpdateUserResponse], error) {
	authMeta := s.app.AuthService().GetMetadata(getAccessToken(r))
	user, err := s.app.UsersService().Update(ctx, &authMeta, &users.UpdateParams{
		Name: r.Msg.Name,
		Dob:  nil,
	})

	if err != nil {
		return nil, serviceErrorHandler(err)
	}
	res := connect.NewResponse(&usersv1.UpdateUserResponse{
		User: mapUser(&user),
	})
	return res, nil
}

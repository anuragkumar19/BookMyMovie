package api

import (
	"context"

	"bookmymovie.app/bookmymovie"
	usersv1 "bookmymovie.app/bookmymovie/api/gen/users/v1"
	"bookmymovie.app/bookmymovie/api/gen/users/v1/usersv1connect"
	"connectrpc.com/connect"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	var dob *date.Date
	if user.Dob.Valid {
		dob = &date.Date{
			Year:  int32(user.Dob.Time.Year()),
			Month: int32(user.Dob.Time.Month()),
			Day:   int32(user.Dob.Time.Day()),
		}
	}
	res := connect.NewResponse(&usersv1.GetLoggedInUserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Version:   user.Version,
		Role:      mapRole(user.Role),
		Dob:       dob,
		CreatedAt: timestamppb.New(user.CreatedAt.Time),
	})
	return res, nil
}

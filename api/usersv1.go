package api

import (
	"context"

	usersv1 "bookmymovie.app/bookmymovie/api/gen/users/v1"
	"bookmymovie.app/bookmymovie/services/users"
	"connectrpc.com/connect"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type usersV1Service struct {
	// usersv1connect.UnimplementedUsersServiceHandler

	users *users.Users
}

func (s *usersV1Service) GetUserInfo(ctx context.Context, r *connect.Request[usersv1.GetUserInfoRequest]) (*connect.Response[usersv1.GetUserInfoResponse], error) {
	info, err := s.users.GetUserInfo(ctx, r.Header().Get("Authorization"))
	if err != nil {
		return nil, err //TODO:
	}
	var dob *date.Date
	if info.Dob.Valid {
		dob = &date.Date{
			Year:  int32(info.Dob.Time.Year()),
			Month: int32(info.Dob.Time.Month()),
			Day:   int32(info.Dob.Time.Day()),
		}
	}
	res := connect.NewResponse(&usersv1.GetUserInfoResponse{
		Id:        info.ID,
		Name:      info.Name,
		Email:     info.Email,
		Version:   info.Version,
		Role:      mapRole(info.Role),
		Dob:       dob,
		CreatedAt: timestamppb.New(info.CreatedAt.Time),
	})
	return res, nil
}

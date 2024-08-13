package api

import (
	"context"

	authv1 "bookmymovie.app/bookmymovie/gen/auth/v1"
	"bookmymovie.app/bookmymovie/services/auth"
	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type authV1Service struct {
	// authv1connect.UnimplementedAuthServiceHandler
	auth *auth.Auth
}

func (s *authV1Service) RequestLoginOTP(ctx context.Context, r *connect.Request[authv1.RequestLoginOTPRequest]) (*connect.Response[authv1.RequestLoginOTPResponse], error) {
	token, err := s.auth.RequestLoginOTP(ctx, &auth.RequestLoginOTPParams{
		Email: r.Msg.Email,
	})
	if err != nil {
		return nil, err //TODO:
	}

	res := connect.NewResponse(&authv1.RequestLoginOTPResponse{
		LoginToken: token,
	})

	return res, nil
}

func (s *authV1Service) Login(ctx context.Context, r *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
	tokens, err := s.auth.Login(ctx, &auth.LoginParams{
		Token:     r.Msg.LoginToken,
		OTP:       r.Msg.Otp,
		UserAgent: r.Header().Get("User-Agent"),
	})
	if err != nil {
		return nil, err //TODO:
	}

	res := connect.NewResponse(&authv1.LoginResponse{
		AccessToken:       tokens.AccessToken,
		RefreshToken:      tokens.RefreshToken,
		AccessTokenExpiry: timestamppb.New(tokens.AccessTokenExpiry),
	})

	return res, nil
}

func (s *authV1Service) RefreshAccessToken(ctx context.Context, r *connect.Request[authv1.RefreshAccessTokenRequest]) (*connect.Response[authv1.RefreshAccessTokenResponse], error) {
	tokens, err := s.auth.RefreshAccessToken(ctx, r.Header().Get("Authorization"))
	if err != nil {
		return nil, err //TODO:
	}

	res := connect.NewResponse(&authv1.RefreshAccessTokenResponse{
		AccessToken:       tokens.AccessToken,
		AccessTokenExpiry: timestamppb.New(tokens.AccessTokenExpiry),
	})

	return res, nil
}

func (s *authV1Service) Logout(ctx context.Context, r *connect.Request[authv1.LogoutRequest]) (*connect.Response[authv1.LogoutResponse], error) {
	err := s.auth.Logout(ctx, r.Header().Get("Authorization"))
	if err != nil {
		return nil, err //TODO:
	}
	res := connect.NewResponse(&authv1.LogoutResponse{})
	return res, nil
}

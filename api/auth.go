package api

import (
	"context"
	"strings"

	authv1 "bookmymovie.app/bookmymovie/api/gen/auth/v1"
	"bookmymovie.app/bookmymovie/services/auth"
	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type authService struct {
	// authv1connect.UnimplementedAuthServiceHandler
	auth *auth.Auth
}

func (s *authService) RequestLoginOTP(ctx context.Context, r *connect.Request[authv1.RequestLoginOTPRequest]) (*connect.Response[authv1.RequestLoginOTPResponse], error) {
	token, err := s.auth.RequestLoginOTP(ctx, &auth.RequestLoginOTPParams{
		Email: r.Msg.GetEmail(),
	})
	if err != nil {
		return nil, err // TODO:
	}

	res := connect.NewResponse(&authv1.RequestLoginOTPResponse{
		LoginToken: token,
	})

	return res, nil
}

func (s *authService) Login(ctx context.Context, r *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
	tokens, err := s.auth.Login(ctx, &auth.LoginParams{
		Token:     r.Msg.GetLoginToken(),
		OTP:       r.Msg.GetOtp(),
		UserAgent: r.Header().Get("User-Agent"),
	})
	if err != nil {
		return nil, err // TODO:
	}

	res := connect.NewResponse(&authv1.LoginResponse{
		AccessToken:       tokens.AccessToken,
		RefreshToken:      tokens.RefreshToken,
		AccessTokenExpiry: timestamppb.New(tokens.AccessTokenExpiry),
	})

	return res, nil
}

func (s *authService) RefreshAccessToken(ctx context.Context, r *connect.Request[authv1.RefreshAccessTokenRequest]) (*connect.Response[authv1.RefreshAccessTokenResponse], error) {
	tokens, err := s.auth.RefreshAccessToken(ctx, r.Header().Get("Authorization"))
	if err != nil {
		return nil, err // TODO:
	}

	res := connect.NewResponse(&authv1.RefreshAccessTokenResponse{
		AccessToken:       tokens.AccessToken,
		AccessTokenExpiry: timestamppb.New(tokens.AccessTokenExpiry),
	})

	return res, nil
}

func (s *authService) Logout(ctx context.Context, r *connect.Request[authv1.LogoutRequest]) (*connect.Response[authv1.LogoutResponse], error) {
	authMeta := s.auth.GetMetadata(getAccessToken(r))
	if err := s.auth.Logout(ctx, &authMeta); err != nil {
		return nil, err // TODO:
	}
	res := connect.NewResponse(&authv1.LogoutResponse{})
	return res, nil
}

func getAccessToken[T any](r *connect.Request[T]) string {
	h := r.Header().Get("Authorization")

	hs := strings.Split(h, " ")
	if len(hs) != 2 {
		return ""
	}
	if hs[0] != "Bearer" {
		return ""
	}
	return hs[1]
}

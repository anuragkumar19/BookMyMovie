package api

import (
	"net/http"

	"bookmymovie.app/bookmymovie"
	"bookmymovie.app/bookmymovie/api/gen/auth/v1/authv1connect"
	"bookmymovie.app/bookmymovie/api/gen/users/v1/usersv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Api struct {
	app *bookmymovie.Application

	authV1Service  *authV1Service
	usersV1Service *usersV1Service
}

func New(app *bookmymovie.Application) Api {
	return Api{
		app:            app,
		authV1Service:  &authV1Service{auth: app.AuthService()},
		usersV1Service: &usersV1Service{users: app.UsersService()},
	}
}

func (api *Api) Run() {
	mux := http.NewServeMux()

	{
		path, handler := authv1connect.NewAuthServiceHandler(api.authV1Service)
		mux.Handle(path, handler)
	}
	{
		path, handler := usersv1connect.NewUsersServiceHandler(api.usersV1Service)
		mux.Handle(path, handler)
	}

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

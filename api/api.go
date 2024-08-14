package api

import (
	"net/http"

	"bookmymovie.app/bookmymovie"
	"bookmymovie.app/bookmymovie/api/gen/auth/v1/authv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Api struct {
	app *bookmymovie.Application

	authV1Service *authV1Service
}

func New(app *bookmymovie.Application) Api {
	return Api{
		app:           app,
		authV1Service: &authV1Service{auth: app.AuthService()},
	}
}

func (api *Api) Run() {
	mux := http.NewServeMux()

	path, handler := authv1connect.NewAuthServiceHandler(api.authV1Service)
	mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

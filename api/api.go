package api

import (
	"net/http"

	"bookmymovie.app/bookmymovie"
	"bookmymovie.app/bookmymovie/api/gen/auth/v1/authv1connect"
	"bookmymovie.app/bookmymovie/api/gen/movies/v1/moviesv1connect"
	"bookmymovie.app/bookmymovie/api/gen/users/v1/usersv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Api struct {
	app *bookmymovie.Application

	authService            *authService
	usersService           *usersService
	moviesLanguagesService *moviesLanguagesService
	moviesFormatsService   *moviesFormatsService
	moviesGenresService    *moviesGenresService
}

func New(app *bookmymovie.Application) Api {
	return Api{
		app:          app,
		authService:  &authService{auth: app.AuthService()},
		usersService: &usersService{users: app.UsersService()},
	}
}

func (api *Api) Run() {
	mux := http.NewServeMux()

	{
		path, handler := authv1connect.NewAuthServiceHandler(api.authService)
		mux.Handle(path, handler)
	}
	{
		path, handler := usersv1connect.NewUsersServiceHandler(api.usersService)
		mux.Handle(path, handler)
	}
	{
		path, handler := moviesv1connect.NewMoviesLanguagesServiceHandler(api.moviesLanguagesService)
		mux.Handle(path, handler)
	}
	{
		path, handler := moviesv1connect.NewMoviesGenresServiceHandler(api.moviesGenresService)
		mux.Handle(path, handler)
	}
	{
		path, handler := moviesv1connect.NewMoviesFormatsServiceHandler(api.moviesFormatsService)
		mux.Handle(path, handler)
	}

	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

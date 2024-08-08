package main

import (
	"bookmymovie.app/bookmymovie"
	"bookmymovie.app/bookmymovie/server"
)

func main() {
	app := bookmymovie.New()

	srv := server.New()

	srv.RegisterService(app.AuthService())
	srv.TempServe()
}

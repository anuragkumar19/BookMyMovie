package main

import (
	"bookmymovie.app/bookmymovie"
	"bookmymovie.app/bookmymovie/api"
)

func main() {
	app := bookmymovie.New()

	apiServer := api.New(&app)

	apiServer.Run()

}

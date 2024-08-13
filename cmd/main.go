package main

import (
	"bookmymovie.app/bookmymovie"
	"bookmymovie.app/bookmymovie/api"
)

func main() {
	app := bookmymovie.New()

	api := api.New(&app)

	api.Run()

}

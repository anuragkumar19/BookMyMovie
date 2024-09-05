package main

import (
	"bookmymovie.app/bookmymovie"
	"bookmymovie.app/bookmymovie/api"
)

func main() {
	appConf := bookmymovie.DefaultConfig()

	app := bookmymovie.New(&appConf)

	apiServer := api.New(&app)

	apiServer.Run()

}

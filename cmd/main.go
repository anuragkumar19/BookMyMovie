package main

import (
	"flag"
	"os"

	"bookmymovie.app/bookmymovie"
	"bookmymovie.app/bookmymovie/api"
	"github.com/rs/zerolog"
)

func main() {
	logLevel := zerolog.InfoLevel

	logLevelFromEnv := os.Getenv("LOG_LEVEL")
	logLevelStr := flag.String("log-level", logLevelFromEnv, "Logging level")

	if *logLevelStr != "" {
		ll, err := zerolog.ParseLevel(*logLevelStr)
		if nil == err {
			logLevel = ll
		}
	}

	logger := zerolog.New(os.Stdout).Level(logLevel).With().Timestamp().Logger()

	config := bookmymovie.DefaultConfig()
	if err := config.ParseFromEnvVars(); err != nil {
		logger.Fatal().Err(err).Msg("failed to load app config from env vars")
	}
	if err := config.ParseFromCLIFlags(); err != nil {
		logger.Fatal().Err(err).Msg("failed to load app config from cli flags")
	}
	app, err := bookmymovie.New(&config, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create instance on bookmymovie.Application")
	}

	apiServer := api.New(&app)
	if err := apiServer.Run(); err != nil {
		logger.Fatal().Err(err).Msg("")
	}
}

package bookmymovie

import (
	"context"
	"os"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"bookmymovie.app/bookmymovie/services/auth"
	"github.com/rs/zerolog"
)

type Application struct {
	db     *database.Database
	mailer *mailer.Mailer

	authService *auth.Auth
}

func New() Application {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	conf := newConfig()
	if err := conf.parse(); err != nil {
		logger.Fatal().Err(err).Msg("failed to parse config")
	}
	if err := conf.validate(); err != nil {
		logger.Fatal().Err(err).Msg("failed to validate config")
	}

	logger = logger.Level(conf.logLevel)

	db := database.NewDatabase(conf.database, &logger)
	m := mailer.New(conf.mailer, &logger)
	authService := auth.New(conf.auth, &logger, &db, &m)

	return Application{
		db:          &db,
		mailer:      &m,
		authService: &authService,
	}
}

func (app *Application) AuthService() *auth.Auth {
	return app.authService
}

func (app *Application) Shutdown(_ context.Context) error {
	app.mailer.Wait()
	return nil
}

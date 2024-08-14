package bookmymovie

import (
	"context"
	"os"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/services/users"
	"github.com/rs/zerolog"
)

type Application struct {
	db     *database.Database
	mailer *mailer.Mailer

	authService  *auth.Auth
	usersService *users.Users
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
	usersService := users.New(&logger, &db, &m, &authService)

	return Application{
		db:           &db,
		mailer:       &m,
		authService:  &authService,
		usersService: &usersService,
	}
}

func (app *Application) AuthService() *auth.Auth {
	return app.authService
}

func (app *Application) UsersService() *users.Users {
	return app.usersService
}

func (app *Application) Shutdown(_ context.Context) error {
	app.mailer.Wait()
	return nil
}

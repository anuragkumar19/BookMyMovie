package bookmymovie

import (
	"context"
	"os"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/services/users"
	"bookmymovie.app/bookmymovie/storage"
	"github.com/rs/zerolog"
)

type Application struct {
	logger  *zerolog.Logger
	db      *database.Database
	storage *storage.Storage
	mailer  *mailer.Mailer

	authService  *auth.Auth
	usersService *users.Users
}

func New(conf *Config) Application {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if err := conf.Validate(); err != nil {
		logger.Fatal().Err(err).Msg("failed to validate config")
	}

	logger = logger.Level(conf.LogLevel)

	db := database.NewDatabase(conf.Database, &logger)
	s := storage.New(conf.Storage, &logger)
	m := mailer.New(conf.Mailer, &logger)
	authService := auth.New(conf.Auth, &logger, &db, &m)
	usersService := users.New(&logger, &db, &m, &authService)

	return Application{
		logger:       &logger,
		db:           &db,
		storage:      &s,
		mailer:       &m,
		authService:  &authService,
		usersService: &usersService,
	}
}

func (app *Application) Logger() *zerolog.Logger {
	return app.logger
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

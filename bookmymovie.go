package bookmymovie

import (
	"context"
	"errors"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/mailer"
	"bookmymovie.app/bookmymovie/services/auth"
	"bookmymovie.app/bookmymovie/services/movies"
	"bookmymovie.app/bookmymovie/services/users"
	"bookmymovie.app/bookmymovie/storage"
	"github.com/rs/zerolog"
)

type Application struct {
	logger  *zerolog.Logger
	db      *database.Database
	storage *storage.Storage
	mailer  *mailer.Mailer

	authService   *auth.Auth
	usersService  *users.Users
	moviesService *movies.Movies
}

func New(config *Config, logger *zerolog.Logger) (Application, error) {
	db, err := database.NewDatabase(config.Database, logger)
	if err != nil {
		return Application{}, errors.Join(errors.New("failed to create instance of database.Database"), err)
	}
	s, err := storage.New(config.Storage, logger)
	if err != nil {
		return Application{}, errors.Join(errors.New("failed to create instance of storage.Storage"), err)
	}
	m, err := mailer.New(config.Mailer, logger)
	if err != nil {
		return Application{}, errors.Join(errors.New("failed to create instance of mailer.Mailer"), err)
	}
	authService, err := auth.New(config.Auth, logger, &db, &m)
	if err != nil {
		return Application{}, errors.Join(errors.New("failed to create instance of auth.Auth"), err)
	}

	usersService := users.New(logger, &db, &m, &authService)

	moviesService, err := movies.New(logger, &db, &authService)
	if err != nil {
		return Application{}, errors.Join(errors.New("failed to create instance of movies.Movies"), err)
	}
	return Application{
		logger:        logger,
		db:            &db,
		storage:       &s,
		mailer:        &m,
		authService:   &authService,
		usersService:  &usersService,
		moviesService: &moviesService,
	}, nil
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

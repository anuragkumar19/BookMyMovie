// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"context"
)

type Querier interface {
	AttemptLoginToken(ctx context.Context, arg *AttemptLoginTokenParams) error
	CreateLoginToken(ctx context.Context, arg *CreateLoginTokenParams) error
	CreateMoviesFormat(ctx context.Context, arg *CreateMoviesFormatParams) (MoviesFormat, error)
	CreateMoviesGenre(ctx context.Context, arg *CreateMoviesGenreParams) (MoviesGenre, error)
	CreateMoviesLanguage(ctx context.Context, arg *CreateMoviesLanguageParams) (MoviesLanguage, error)
	CreateMoviesPerson(ctx context.Context, arg *CreateMoviesPersonParams) (int64, error)
	CreateRefreshToken(ctx context.Context, arg *CreateRefreshTokenParams) (RefreshToken, error)
	CreateRegularUser(ctx context.Context, email string) (int64, error)
	DeleteExpiredLoginTokens(ctx context.Context) error
	DeleteExpiredRefreshTokens(ctx context.Context) error
	DeleteLoginToken(ctx context.Context, token string) error
	DeleteMoviesFormat(ctx context.Context, id int64) error
	DeleteMoviesGenre(ctx context.Context, id int64) error
	DeleteMoviesLanguage(ctx context.Context, id int64) error
	DeleteRefreshToken(ctx context.Context, id int64) error
	FindLoginToken(ctx context.Context, token string) (FindLoginTokenRow, error)
	FindRefreshToken(ctx context.Context, token string) (RefreshToken, error)
	FindUserByEmail(ctx context.Context, email string) (FindUserByEmailRow, error)
	FindUserById(ctx context.Context, id int64) (FindUserByIdRow, error)
	GetAllMoviesFormats(ctx context.Context) ([]MoviesFormat, error)
	GetAllMoviesGenres(ctx context.Context) ([]MoviesGenre, error)
	GetAllMoviesLanguages(ctx context.Context) ([]MoviesLanguage, error)
	GetMoviesFormatByID(ctx context.Context, id int64) (MoviesFormat, error)
	GetMoviesGenreByID(ctx context.Context, id int64) (MoviesGenre, error)
	GetMoviesLanguageByID(ctx context.Context, id int64) (MoviesLanguage, error)
	UpdateMoviesFormat(ctx context.Context, arg *UpdateMoviesFormatParams) error
	UpdateMoviesGenre(ctx context.Context, arg *UpdateMoviesGenreParams) error
	UpdateMoviesLanguage(ctx context.Context, arg *UpdateMoviesLanguageParams) error
	UpdateUserLoginFields(ctx context.Context, arg *UpdateUserLoginFieldsParams) error
}

var _ Querier = (*Queries)(nil)

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
	CreateRegularUser(ctx context.Context, email string) (int64, error)
	DeleteExpiredTokens(ctx context.Context) error
	DeleteLoginToken(ctx context.Context, token string) error
	FindLoginToken(ctx context.Context, token string) (FindLoginTokenRow, error)
	FindUserByEmail(ctx context.Context, email string) (FindUserByEmailRow, error)
	UpdateUserLoginFields(ctx context.Context, arg *UpdateUserLoginFieldsParams) error
}

var _ Querier = (*Queries)(nil)

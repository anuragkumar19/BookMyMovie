// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: refresh_tokens.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRefreshToken = `-- name: CreateRefreshToken :one
INSERT INTO
    "refresh_tokens" ("token", "created_at", "user_id", "user_role", "expire_at", "user_agent")
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING
    "id",
    "token",
    "created_at",
    "user_id",
    "user_role",
    "expire_at",
    "user_agent"
`

type CreateRefreshTokenParams struct {
	Token     string
	CreatedAt pgtype.Timestamptz
	UserID    int64
	UserRole  Roles
	ExpireAt  pgtype.Timestamptz
	UserAgent string
}

func (q *Queries) CreateRefreshToken(ctx context.Context, arg *CreateRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRow(ctx, createRefreshToken,
		arg.Token,
		arg.CreatedAt,
		arg.UserID,
		arg.UserRole,
		arg.ExpireAt,
		arg.UserAgent,
	)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.CreatedAt,
		&i.UserID,
		&i.UserRole,
		&i.ExpireAt,
		&i.UserAgent,
	)
	return i, err
}

const deleteExpiredRefreshTokens = `-- name: DeleteExpiredRefreshTokens :exec
DELETE FROM "refresh_tokens"
WHERE
    "expire_at" < NOW()
`

func (q *Queries) DeleteExpiredRefreshTokens(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteExpiredRefreshTokens)
	return err
}

const deleteRefreshToken = `-- name: DeleteRefreshToken :exec
DELETE FROM "refresh_tokens"
WHERE
    "id" = $1
`

func (q *Queries) DeleteRefreshToken(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteRefreshToken, id)
	return err
}

const findRefreshToken = `-- name: FindRefreshToken :one
SELECT
    "id",
    "token",
    "created_at",
    "user_id",
    "user_role",
    "expire_at",
    "user_agent"
FROM
    "refresh_tokens"
WHERE
    "token" = $1
`

func (q *Queries) FindRefreshToken(ctx context.Context, token string) (RefreshToken, error) {
	row := q.db.QueryRow(ctx, findRefreshToken, token)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.Token,
		&i.CreatedAt,
		&i.UserID,
		&i.UserRole,
		&i.ExpireAt,
		&i.UserAgent,
	)
	return i, err
}

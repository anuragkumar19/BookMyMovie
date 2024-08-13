// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: login_tokens.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const attemptLoginToken = `-- name: AttemptLoginToken :exec
UPDATE "login_tokens"
SET
    last_attempt_at = $1,
    total_attempts = $2
WHERE
    "token" = $3
    AND "version" = $4
`

type AttemptLoginTokenParams struct {
	LastAttemptAt pgtype.Timestamptz
	TotalAttempts int32
	Token         string
	Version       int32
}

func (q *Queries) AttemptLoginToken(ctx context.Context, arg *AttemptLoginTokenParams) error {
	_, err := q.db.Exec(ctx, attemptLoginToken,
		arg.LastAttemptAt,
		arg.TotalAttempts,
		arg.Token,
		arg.Version,
	)
	return err
}

const createLoginToken = `-- name: CreateLoginToken :exec
INSERT INTO
    "login_tokens" (token, otp, created_at, expire_at, user_id)
VALUES
    ($1, $2, $3, $4, $5)
`

type CreateLoginTokenParams struct {
	Token     string
	Otp       string
	CreatedAt pgtype.Timestamptz
	ExpireAt  pgtype.Timestamptz
	UserID    int64
}

func (q *Queries) CreateLoginToken(ctx context.Context, arg *CreateLoginTokenParams) error {
	_, err := q.db.Exec(ctx, createLoginToken,
		arg.Token,
		arg.Otp,
		arg.CreatedAt,
		arg.ExpireAt,
		arg.UserID,
	)
	return err
}

const deleteExpiredLoginTokens = `-- name: DeleteExpiredLoginTokens :exec
DELETE FROM "login_tokens"
WHERE
    expire_at < NOW()
`

func (q *Queries) DeleteExpiredLoginTokens(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteExpiredLoginTokens)
	return err
}

const deleteLoginToken = `-- name: DeleteLoginToken :exec
DELETE FROM "login_tokens"
WHERE
    token = $1
`

func (q *Queries) DeleteLoginToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, deleteLoginToken, token)
	return err
}

const findLoginToken = `-- name: FindLoginToken :one
SELECT
    login_tokens.token,
    login_tokens.otp,
    login_tokens."version",
    login_tokens.user_id,
    login_tokens.created_at,
    login_tokens.expire_at,
    login_tokens.last_attempt_at,
    login_tokens.total_attempts,
    users.role AS user_role
FROM
    "login_tokens"
    INNER JOIN users ON login_tokens.user_id = users.id
WHERE
    token = $1
`

type FindLoginTokenRow struct {
	Token         string
	Otp           string
	Version       int32
	UserID        int64
	CreatedAt     pgtype.Timestamptz
	ExpireAt      pgtype.Timestamptz
	LastAttemptAt pgtype.Timestamptz
	TotalAttempts int32
	UserRole      Roles
}

func (q *Queries) FindLoginToken(ctx context.Context, token string) (FindLoginTokenRow, error) {
	row := q.db.QueryRow(ctx, findLoginToken, token)
	var i FindLoginTokenRow
	err := row.Scan(
		&i.Token,
		&i.Otp,
		&i.Version,
		&i.UserID,
		&i.CreatedAt,
		&i.ExpireAt,
		&i.LastAttemptAt,
		&i.TotalAttempts,
		&i.UserRole,
	)
	return i, err
}

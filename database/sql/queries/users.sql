-- name: FindUserByEmail :one
SELECT
    "id",
    "email",
    "version",
    "last_login_token_sent_at",
    "total_login_tokens_sent"
FROM
    "users"
WHERE
    "email" = $1;

-- name: CreateRegularUser :one
INSERT INTO
    "users" ("email", "role")
VALUES
    ($1, 'regular_user')
RETURNING
    id;

-- name: UpdateUserLoginFields :exec
UPDATE "users"
SET
    "last_login_token_sent_at" = $1,
    "total_login_tokens_sent" = $2
WHERE
    "id" = $3
    AND "version" = $4;
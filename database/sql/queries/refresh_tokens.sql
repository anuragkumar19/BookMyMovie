-- name: CreateRefreshToken :one
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
    "user_agent";

-- name: FindRefreshToken :one
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
    "token" = $1;

-- name: DeleteRefreshToken :exec
DELETE FROM "refresh_tokens"
WHERE
    "id" = $1;

-- name: DeleteExpiredRefreshTokens :exec
DELETE FROM "refresh_tokens"
WHERE
    "expire_at" < NOW();
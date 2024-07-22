-- name: CreateLoginToken :exec
INSERT INTO
    "login_tokens" (token, otp, created_at, expire_at, user_id)
VALUES
    ($1, $2, $3, $4, $5);

-- name: FindLoginToken :one
SELECT
    token,
    otp,
    "version",
    created_at,
    expire_at,
    last_attempt_at,
    total_attempts
FROM
    "login_tokens"
WHERE
    token = $1;

-- name: AttemptLoginToken :exec
UPDATE "login_tokens"
SET
    last_attempt_at = $1,
    total_attempts = $2
WHERE
    "token" = $3
    AND "version" = $4;

-- name: DeleteLoginToken :exec
DELETE FROM "login_tokens"
WHERE
    token = $1;

-- name: DeleteExpiredTokens :exec
DELETE FROM "login_tokens"
WHERE
    expire_at < NOW();
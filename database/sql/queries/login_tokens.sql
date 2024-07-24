-- name: CreateLoginToken :exec
INSERT INTO
    "login_tokens" (token, otp, created_at, expire_at, user_id)
VALUES
    ($1, $2, $3, $4, $5);

-- name: FindLoginToken :one
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

-- name: DeleteExpiredLoginTokens :exec
DELETE FROM "login_tokens"
WHERE
    expire_at < NOW();
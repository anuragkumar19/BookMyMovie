-- name: CreateMoviesLanguage :one
INSERT INTO
    "movies_languages" ("slug", "display_name", "english_name")
VALUES
    ($1, $2, $3)
RETURNING
    *;

-- name: UpdateMoviesLanguage :exec
UPDATE "movies_languages"
SET
    "slug" = $1,
    "display_name" = $2,
    "english_name" = $3
WHERE
    "id" = $4
    AND "version" = $5;

-- name: GetAllMoviesLanguages :many
SELECT
    *
FROM
    "movies_languages";

-- name: GetMoviesLanguageByID :one
SELECT
    *
FROM
    "movies_languages"
WHERE
    "id" = $1;

-- name: DeleteMoviesLanguage :exec
DELETE FROM "movies_languages"
WHERE
    id = $1;
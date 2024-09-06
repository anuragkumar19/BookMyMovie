-- name: CreateMoviesLanguage :one
INSERT INTO
    "movies_languages" ("slug", "display_name", "english_name")
VALUES
    ($1, $2, $3)
RETURNING
    *;

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
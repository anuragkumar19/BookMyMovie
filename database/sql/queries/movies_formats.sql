-- name: CreateMoviesFormat :one
INSERT INTO
    "movies_formats" ("slug", "display_name", "about")
VALUES
    ($1, $2, $3)
RETURNING
    *;

-- name: GetAllMoviesFormats :many
SELECT
    *
FROM
    "movies_formats";

-- name: GetMoviesFormatByID :one
SELECT
    *
FROM
    "movies_formats"
WHERE
    "id" = $1;

-- name: DeleteMoviesFormat :exec
DELETE FROM "movies_formats"
WHERE
    id = $1;
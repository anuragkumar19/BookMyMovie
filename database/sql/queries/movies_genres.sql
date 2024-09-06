-- name: CreateMoviesGenre :one
INSERT INTO
    "movies_genres" ("slug", "display_name", "about")
VALUES
    ($1, $2, $3)
RETURNING
    *;

-- name: GetAllMoviesGenres :many
SELECT
    *
FROM
    "movies_genres";

-- name: GetMoviesGenreByID :one
SELECT
    *
FROM
    "movies_genres"
WHERE
    "id" = $1;

-- name: DeleteMoviesGenre :exec
DELETE FROM "movies_genres"
WHERE
    id = $1;
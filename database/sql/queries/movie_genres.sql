-- name: CreateGenre :one
INSERT INTO
    "movie_genres" ("id", "display_name", "about")
VALUES
    ($1, $2, $3)
RETURNING
    *;

-- name: GetAllGenres :many
SELECT
    *
FROM
    "movie_genres";

-- name: GetGenreByID :one
SELECT
    *
FROM
    "movie_genres"
WHERE
    "id" = $1;

-- name: UpdateGenre :exec
UPDATE "movie_genres"
SET
    "display_name" = $1,
    "about" = $2
WHERE
    id = $3;

-- name: DeleteGenre :exec
DELETE FROM "movie_genres"
WHERE
    id = $1;
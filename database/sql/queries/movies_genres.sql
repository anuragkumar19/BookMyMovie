-- name: CreateMoviesGenre :one
INSERT INTO
    "movies_genres" ("slug", "display_name", "about")
VALUES
    ($1, $2, $3)
RETURNING
    *;

-- name: CheckIfAnyMoviesGenresJoinExist :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            "movies_genres_join"
        WHERE
            "movies_genre_id" = $1
    );

-- name: UpdateMoviesGenre :exec
UPDATE "movies_genres"
SET
    "slug" = $1,
    "display_name" = $2,
    "about" = $3
WHERE
    "id" = $4
    AND "version" = $5;

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
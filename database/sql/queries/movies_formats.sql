-- name: CreateMoviesFormat :one
INSERT INTO
    "movies_formats" ("slug", "display_name", "about")
VALUES
    ($1, $2, $3)
RETURNING
    *;

-- name: CheckIfAnyMoviesAvailableFormatsExist :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            "movies_available_formats"
        WHERE
            "movies_format_id" = $1
    );

-- name: UpdateMoviesFormat :exec
UPDATE "movies_formats"
SET
    "slug" = $1,
    "display_name" = $2,
    "about" = $3,
    "version" = "version" + 1
WHERE
    "id" = $4
    AND "version" = $5;

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
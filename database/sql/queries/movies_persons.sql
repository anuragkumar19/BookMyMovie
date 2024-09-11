-- name: CreateMoviesPerson :one
INSERT INTO
    "movies_persons" ("name", "slug", "nicknames", "profile_picture", "occupations", "dob", "about")
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    "id";

-- name: GetMoviesPerson :one
SELECT
    "id",
    "name",
    "slug",
    "nicknames",
    "profile_picture",
    "occupations",
    "dob",
    "about",
    "version"
FROM
    "movies_persons"
WHERE
    "id" = $1
    AND "is_deleted" = FALSE;

-- name: ListMoviesPerson :many
SELECT
    "id",
    "name",
    "slug",
    "nicknames",
    "profile_picture",
    "occupations",
    "dob",
    "about" "version"
FROM
    "movies_persons"
WHERE
    (
        "created_at" < $1
        AND "is_deleted" = FALSE
    )
    OR (
        "deleted_at" > $1
        AND "is_deleted" = TRUE
    )
ORDER BY
    "created_at" DESC
LIMIT
    $2
OFFSET
    $3;

-- name: UpdateMoviesPerson :exec
UPDATE "movies_persons"
SET
    "name" = $1,
    "slug" = $2,
    "nicknames" = $3,
    "profile_picture" = $4,
    "occupations" = $5,
    "dob" = $6,
    "about" = $7,
    "version" = "version" + 1
WHERE
    "id" = $8
    AND "version" = $9
    AND "is_deleted" = FALSE;

-- name: DeleteMoviesPerson :exec
UPDATE "movies_persons"
SET
    "is_deleted" = TRUE,
    "version" = "version" + 1
WHERE
    "id" = $1
    AND "version" = $2
    AND "is_deleted" = FALSE;
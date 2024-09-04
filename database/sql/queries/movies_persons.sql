-- name: CreateMoviesPerson :one
INSERT INTO
    "movies_persons" (
        "name",
        "slug",
        "nicknames",
        "profile_picture",
        "occupations",
        "dob",
        "about",
        "imdb_id",
        "imdb_last_synced_at"
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING
    "id";
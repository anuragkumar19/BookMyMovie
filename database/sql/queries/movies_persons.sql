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
    "about",
    "version"
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

-- name: SearchMoviesPerson :many
SELECT
    "movies_persons"."id",
    "movies_persons"."name",
    "movies_persons"."slug",
    "movies_persons"."nicknames",
    "movies_persons"."profile_picture",
    "movies_persons"."occupations",
    "movies_persons"."dob",
    "movies_persons"."about",
    "movies_persons"."version"
FROM
    "movies_persons",
    COALESCE(text_array_to_text ("movies_persons"."nicknames"), '') nicknames_text,
    COALESCE(text_array_to_text ("movies_persons"."occupations"), '') occupations_text,
    TO_TSVECTOR(
        'english',
        "movies_persons"."name" || nicknames_text || occupations_text || "movies_persons"."about"
    ) DOCUMENT,
    TO_TSQUERY($1) query,
    NULLIF(TS_RANK(TO_TSVECTOR("movies_persons"."name"), query), 0) rank_name,
    NULLIF(TS_RANK(TO_TSVECTOR(nicknames_text), query), 0) rank_nicknames,
    NULLIF(TS_RANK(TO_TSVECTOR(occupations_text), query), 0) rank_occupations,
    NULLIF(TS_RANK(TO_TSVECTOR("movies_persons"."about"), query), 0) rank_about
WHERE
    (
        (
            "created_at" < $2
            AND "is_deleted" = FALSE
        )
        OR (
            "deleted_at" > $2
            AND "is_deleted" = TRUE
        )
    )
    AND (query @@ DOCUMENT)
ORDER BY
    rank_name,
    rank_nicknames,
    rank_occupations,
    rank_about,
    id DESC NULLS LAST
LIMIT
    $3
OFFSET
    $4;

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
    "version" = "version" + 1,
    "is_deleted" = TRUE,
    "deleted_at" = NOW(),
    "name" = '',
    "slug" = '',
    "nicknames" = '{}',
    "profile_picture" = '',
    "occupations" = '{}',
    "dob" = NOW(),
    "about" = ''
WHERE
    "id" = $1
    AND "version" = $2
    AND "is_deleted" = FALSE;
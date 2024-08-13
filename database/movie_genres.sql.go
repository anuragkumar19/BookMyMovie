// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: movie_genres.sql

package database

import (
	"context"
)

const createGenre = `-- name: CreateGenre :one
INSERT INTO
    "movie_genres" ("id", "display_name", "about")
VALUES
    ($1, $2, $3)
RETURNING
    id, created_at, display_name, about
`

type CreateGenreParams struct {
	ID          string
	DisplayName string
	About       string
}

func (q *Queries) CreateGenre(ctx context.Context, arg *CreateGenreParams) (MovieGenre, error) {
	row := q.db.QueryRow(ctx, createGenre, arg.ID, arg.DisplayName, arg.About)
	var i MovieGenre
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.DisplayName,
		&i.About,
	)
	return i, err
}

const deleteGenre = `-- name: DeleteGenre :exec
DELETE FROM "movie_genres"
WHERE
    id = $1
`

func (q *Queries) DeleteGenre(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteGenre, id)
	return err
}

const getAllGenres = `-- name: GetAllGenres :many
SELECT
    id, created_at, display_name, about
FROM
    "movie_genres"
`

func (q *Queries) GetAllGenres(ctx context.Context) ([]MovieGenre, error) {
	rows, err := q.db.Query(ctx, getAllGenres)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MovieGenre
	for rows.Next() {
		var i MovieGenre
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.DisplayName,
			&i.About,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGenreByID = `-- name: GetGenreByID :one
SELECT
    id, created_at, display_name, about
FROM
    "movie_genres"
WHERE
    "id" = $1
`

func (q *Queries) GetGenreByID(ctx context.Context, id string) (MovieGenre, error) {
	row := q.db.QueryRow(ctx, getGenreByID, id)
	var i MovieGenre
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.DisplayName,
		&i.About,
	)
	return i, err
}

const updateGenre = `-- name: UpdateGenre :exec
UPDATE "movie_genres"
SET
    "display_name" = $1,
    "about" = $2
WHERE
    id = $3
`

type UpdateGenreParams struct {
	DisplayName string
	About       string
	ID          string
}

func (q *Queries) UpdateGenre(ctx context.Context, arg *UpdateGenreParams) error {
	_, err := q.db.Exec(ctx, updateGenre, arg.DisplayName, arg.About, arg.ID)
	return err
}
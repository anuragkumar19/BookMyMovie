// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: movies_languages.sql

package database

import (
	"context"
)

const createMoviesLanguage = `-- name: CreateMoviesLanguage :one
INSERT INTO
    "movies_languages" ("id", "display_name", "english_name")
VALUES
    ($1, $2, $3)
RETURNING
    id, created_at, display_name, english_name
`

type CreateMoviesLanguageParams struct {
	ID          string
	DisplayName string
	EnglishName string
}

func (q *Queries) CreateMoviesLanguage(ctx context.Context, arg *CreateMoviesLanguageParams) (MoviesLanguage, error) {
	row := q.db.QueryRow(ctx, createMoviesLanguage, arg.ID, arg.DisplayName, arg.EnglishName)
	var i MoviesLanguage
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.DisplayName,
		&i.EnglishName,
	)
	return i, err
}

const deleteMoviesLanguage = `-- name: DeleteMoviesLanguage :exec
DELETE FROM "movies_languages"
WHERE
    id = $1
`

func (q *Queries) DeleteMoviesLanguage(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteMoviesLanguage, id)
	return err
}

const getAllMoviesLanguages = `-- name: GetAllMoviesLanguages :many
SELECT
    id, created_at, display_name, english_name
FROM
    "movies_languages"
`

func (q *Queries) GetAllMoviesLanguages(ctx context.Context) ([]MoviesLanguage, error) {
	rows, err := q.db.Query(ctx, getAllMoviesLanguages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MoviesLanguage
	for rows.Next() {
		var i MoviesLanguage
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.DisplayName,
			&i.EnglishName,
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

const getMoviesLanguageByID = `-- name: GetMoviesLanguageByID :one
SELECT
    id, created_at, display_name, english_name
FROM
    "movies_languages"
WHERE
    "id" = $1
`

func (q *Queries) GetMoviesLanguageByID(ctx context.Context, id string) (MoviesLanguage, error) {
	row := q.db.QueryRow(ctx, getMoviesLanguageByID, id)
	var i MoviesLanguage
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.DisplayName,
		&i.EnglishName,
	)
	return i, err
}
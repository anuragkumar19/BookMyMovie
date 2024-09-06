-- +goose Up
CREATE TYPE "movies_mpa_ratings" AS ENUM('G', 'PG', 'PG-13', 'R', 'NC-17');

-- +goose Down
DROP TYPE "movies_mpa_ratings";
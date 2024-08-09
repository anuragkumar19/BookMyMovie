-- +goose Up
CREATE TYPE "mpa_ratings" AS ENUM('G', 'PG', 'PG-13', 'R', 'NC-17');

-- +goose Down
DROP TYPE "mpa_ratings";
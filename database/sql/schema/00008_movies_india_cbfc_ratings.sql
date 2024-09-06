-- +goose Up
CREATE TYPE "movies_india_cbfc_ratings" AS ENUM('U', 'U/A', 'A', 'S');

-- +goose Down
DROP TYPE "movies_india_cbfc_ratings";
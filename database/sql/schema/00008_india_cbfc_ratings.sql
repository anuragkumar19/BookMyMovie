-- +goose Up
CREATE TYPE "india_cbfc_ratings" AS ENUM('U', 'U/A', 'A', 'S');

-- +goose Down
DROP TYPE "india_cbfc_ratings";
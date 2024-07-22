-- +goose Up
CREATE TYPE "roles" AS ENUM('regular_user');

-- +goose Down
DROP TYPE "roles";
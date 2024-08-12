-- +goose Up
CREATE TYPE "roles" AS ENUM('regular_user', 'admin');

-- +goose Down
DROP TYPE "roles";
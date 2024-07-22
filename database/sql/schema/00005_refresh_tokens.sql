-- +goose Up
CREATE TABLE
    "refresh_tokens" (
        "token" TEXT PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
        "user_id" bytea NOT NULL,
        "expire_at" TIMESTAMP WITH TIME ZONE NOT NULL
    );

-- +goose Down
DROP TABLE "refresh_tokens";
-- +goose Up
CREATE TABLE
    "refresh_tokens" (
        "id" BIGSERIAL UNIQUE NOT NULL,
        "token" TEXT PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
        "user_id" BIGINT NOT NULL,
        "user_role" roles NOT NULL,
        "expire_at" TIMESTAMP WITH TIME ZONE NOT NULL,
        "user_agent" TEXT NOT NULL,
        CONSTRAINT fk_refresh_tokens_user_id FOREIGN KEY ("user_id") REFERENCES users ("id")
    );

-- +goose Down
DROP TABLE "refresh_tokens";
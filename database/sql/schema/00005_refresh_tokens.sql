-- +goose Up
CREATE TABLE
    "refresh_tokens" (
        "token" BIGSERIAL PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
        "user_id" BIGINT NOT NULL,
        "expire_at" TIMESTAMP WITH TIME ZONE NOT NULL,
        CONSTRAINT fk_user_id FOREIGN KEY ("user_id") REFERENCES users ("id")
    );

-- +goose Down
DROP TABLE "refresh_tokens";
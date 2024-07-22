-- +goose Up
CREATE TABLE
    "login_tokens" (
        "token" TEXT PRIMARY KEY NOT NULL,
        "version" INTEGER DEFAULT 1 NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
        "user_id" bytea NOT NULL,
        "otp" TEXT NOT NULL,
        "expire_at" TIMESTAMP WITH TIME ZONE NOT NULL,
        "last_attempt_at" TIMESTAMP WITH TIME ZONE,
        "total_attempts" INTEGER DEFAULT 0 NOT NULL
    );

-- +goose Down
DROP TABLE "login_tokens";
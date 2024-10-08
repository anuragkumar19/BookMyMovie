-- +goose Up
CREATE TABLE
    "users" (
        "id" BIGSERIAL PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "version" INTEGER DEFAULT 1 NOT NULL,
        "name" TEXT NOT NULL DEFAULT 'User',
        "email" CITEXT NOT NULL UNIQUE,
        "role" roles NOT NULL,
        "dob" DATE,
        "last_login_token_sent_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "total_login_tokens_sent" INTEGER DEFAULT 0 NOT NULL
    );

-- +goose Down
DROP TABLE "users";
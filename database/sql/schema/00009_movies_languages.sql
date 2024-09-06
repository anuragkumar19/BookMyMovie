-- +goose Up
CREATE TABLE
    "movies_languages" (
        "id" BIGSERIAL PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "display_name" TEXT NOT NULL,
        "english_name" TEXT NOT NULL,
        "slug" TEXT NOT NULL
    );

-- +goose Down
DROP TABLE "movies_languages";
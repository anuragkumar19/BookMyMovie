-- +goose Up
CREATE TABLE
    "movie_languages" (
        "id" TEXT PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "display_name" TEXT NOT NULL,
        "english_name" TEXT NOT NULL
    );

-- +goose Down
DROP TABLE "movie_languages";
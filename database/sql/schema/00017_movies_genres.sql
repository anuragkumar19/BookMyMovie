-- +goose Up
CREATE TABLE
    "movies_genres" (
        "id" TEXT PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "display_name" TEXT NOT NULL,
        "about" TEXT NOT NULL
    );

-- +goose Down
DROP TABLE "movies_genres";
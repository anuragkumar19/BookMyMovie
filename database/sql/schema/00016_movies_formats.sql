-- +goose Up
CREATE TABLE
    "movies_formats" (
        "id" TEXT PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "about" TEXT NOT NULL
    );

-- +goose Down
DROP TABLE "movies_formats";
-- +goose Up
CREATE TABLE
    "cities" (
        "id" TEXT PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "display_name" TEXT NOT NULL,
        "state" TEXT NOT NULL,
        "coordinates" TEXT NOT NULL
    );

-- +goose Down
DROP TABLE "cities";
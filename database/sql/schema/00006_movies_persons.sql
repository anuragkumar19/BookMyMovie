-- +goose Up
CREATE TABLE
    "movies_persons" (
        "id" BIGSERIAL PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "version" INTEGER DEFAULT 1 NOT NULL,
        "name" TEXT NOT NULL,
        "slug" TEXT NOT NULL,
        "nicknames" TEXT[] NOT NULL,
        "profile_picture" TEXT NOT NULL,
        "occupations" TEXT[] NOT NULL,
        "dob" DATE,
        "about" TEXT NOT NULL DEFAULT 'No Information available',
        "is_deleted" BOOLEAN NOT NULL DEFAULT FALSE,
        "deleted_at" TIMESTAMP WITH TIME ZONE
    );

-- +goose Down
DROP TABLE "movies_persons";
-- +goose Up
CREATE TABLE
    "movies" (
        "id" BIGSERIAL PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "version" INTEGER DEFAULT 1 NOT NULL,
        "title" TEXT NOT NULL,
        "slug" TEXT NOT NULL,
        "poster" TEXT NOT NULL,
        "running_duration_in_minutes" INTEGER NOT NULL,
        "available_languages" TEXT[] NOT NULL,
        "available_formats" TEXT[] NOT NULL,
        "rating" TEXT NOT NULL,
        "total_rating_count" INTEGER NOT NULL,
        "genres" TEXT[] NOT NULL,
        "release_date" DATE NOT NULL,
        "is_in_cinema" BOOLEAN NOT NULL,
        "india_cbfc_rating" india_cbfc_ratings NOT NULL,
        "mpa_rating" mpa_ratings NOT NULL,
        "about" TEXT NOT NULL DEFAULT 'No Information available',
        "imdb_id" TEXT NOT NULL,
        "imdb_rating" TEXT NOT NULL,
        "imdb_last_synced_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
    );

-- +goose Down
DROP TABLE "movies";
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
        "available_audio_languages" BIGINT[] NOT NULL,
        "available_subtitle_languages" BIGINT[] NOT NULL,
        "available_formats" BIGINT[] NOT NULL,
        "rating" TEXT NOT NULL,
        "total_rating_count" INTEGER NOT NULL,
        "genres" BIGINT[] NOT NULL,
        "release_date" DATE NOT NULL,
        "is_in_cinema" BOOLEAN NOT NULL,
        "india_cbfc_rating" movies_india_cbfc_ratings NOT NULL,
        "mpa_rating" movies_mpa_ratings NOT NULL,
        "about" TEXT NOT NULL DEFAULT 'No Information available'
    );

-- +goose Down
DROP TABLE "movies";
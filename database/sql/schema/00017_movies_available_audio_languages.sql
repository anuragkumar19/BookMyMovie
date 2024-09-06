-- +goose Up
CREATE TABLE
    "movies_available_audio_languages" (
        "movie_id" BIGINT NOT NULL,
        "movies_language_id" BIGINT NOT NULL,
        "index" INTEGER NOT NULL,
        CONSTRAINT fk_movies_available_audio_languages_movies_movie_id FOREIGN KEY ("movie_id") REFERENCES movies ("id"),
        CONSTRAINT fk_movies_available_audio_languages_movies_languages_movies_language_id FOREIGN KEY ("movies_language_id") REFERENCES movies_languages ("id"),
        CONSTRAINT movies_available_audio_languages_pk PRIMARY KEY ("movie_id", "movies_language_id")
    );

-- +goose Down
DROP TABLE "movies_available_audio_languages";
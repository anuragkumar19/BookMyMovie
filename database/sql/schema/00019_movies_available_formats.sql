-- +goose Up
CREATE TABLE
    "movies_available_formats" (
        "movie_id" BIGINT NOT NULL,
        "movies_format_id" BIGINT NOT NULL,
        "index" INTEGER NOT NULL,
        CONSTRAINT fk_movies_available_formats_movies_movie_id FOREIGN KEY ("movie_id") REFERENCES movies ("id"),
        CONSTRAINT fk_movies_available_formats_movies_formats_movies_format_id FOREIGN KEY ("movies_format_id") REFERENCES movies_formats ("id"),
        CONSTRAINT movies_available_formats_pk PRIMARY KEY ("movie_id", "movies_format_id")
    );

-- +goose Down
DROP TABLE "movies_available_formats";
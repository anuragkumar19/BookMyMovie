-- +goose Up
CREATE TABLE
    "movies_genres_join" (
        "movie_id" BIGINT NOT NULL,
        "movies_genre_id" BIGINT NOT NULL,
        "index" INTEGER NOT NULL,
        CONSTRAINT fk_movies_genres_join_movies_movie_id FOREIGN KEY ("movie_id") REFERENCES movies ("id"),
        CONSTRAINT fk_movies_genres_join_movies_genres_movies_genre_id FOREIGN KEY ("movies_genre_id") REFERENCES movies_genres ("id"),
        CONSTRAINT movies_genres_join_pk PRIMARY KEY ("movie_id", "movies_genre_id")
    );

-- +goose Down
DROP TABLE "movies_genres_join";
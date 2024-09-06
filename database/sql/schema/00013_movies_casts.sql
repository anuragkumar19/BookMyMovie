-- +goose Up
CREATE TABLE
    "movies_casts" (
        "movie_id" BIGINT NOT NULL,
        "movies_person_id" BIGINT NOT NULL,
        "index" INTEGER NOT NULL,
        CONSTRAINT fk_movies_casts_movies_movie_id FOREIGN KEY ("movie_id") REFERENCES movies ("id"),
        CONSTRAINT fk_movies_casts_movies_persons_movies_person_id FOREIGN KEY ("movies_person_id") REFERENCES movies_persons ("id"),
        CONSTRAINT movies_casts_pk PRIMARY KEY ("movie_id", "movies_person_id")
    );

-- +goose Down
DROP TABLE "movies_casts";
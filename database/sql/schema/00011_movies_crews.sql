-- +goose Up
CREATE TABLE
    "movies_crews" (
        "movie_id" BIGINT NOT NULL,
        "person_id" BIGINT NOT NULL,
        "index" INTEGER NOT NULL,
        CONSTRAINT fk_movies_crews_movies_movie_id FOREIGN KEY ("movie_id") REFERENCES movies ("id"),
        CONSTRAINT fk_movies_crews_persons_person_id FOREIGN KEY ("person_id") REFERENCES persons ("id"),
        CONSTRAINT movies_crews_pk PRIMARY KEY ("movie_id", "person_id")
    );

-- +goose Down
DROP TABLE "movies_crews";
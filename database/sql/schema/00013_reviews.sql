-- +goose Up
CREATE TABLE
    "reviews" (
        "id" BIGSERIAL PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "version" INTEGER DEFAULT 1 NOT NULL,
        "movie_id" BIGINT NOT NULL,
        "author_id" BIGINT NOT NULL,
        "text" TEXT NOT NULL DEFAULT '',
        "rating" INTEGER NOT NULL,
        CONSTRAINT fk_reviews_movies_movie_id FOREIGN KEY ("movie_id") REFERENCES movies ("id"),
        CONSTRAINT fk_reviews_users_author_id FOREIGN KEY ("author_id") REFERENCES users ("id"),
        CONSTRAINT unq_movie_id_author_id UNIQUE ("movie_id", "author_id")
    );

-- +goose Down
DROP TABLE "reviews";
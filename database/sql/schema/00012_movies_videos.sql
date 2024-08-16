-- +goose Up
CREATE TABLE
    "movies_videos" (
        "id" BIGSERIAL PRIMARY KEY NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
        "version" INTEGER DEFAULT 1 NOT NULL,
        "movie_id" BIGINT NOT NULL,
        "language" TEXT NOT NULL,
        "youtube_link" TEXT NOT NULL,
        "index" INTEGER NOT NULL,
        CONSTRAINT fk_movies_videos_movies_movie_id FOREIGN KEY ("movie_id") REFERENCES movies ("id")
    );

-- +goose Down
DROP TABLE "movies_videos";
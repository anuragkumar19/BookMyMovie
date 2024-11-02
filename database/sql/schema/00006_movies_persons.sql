-- +goose Up
CREATE FUNCTION text_array_to_text (TEXT[]) RETURNS TEXT LANGUAGE SQL IMMUTABLE AS $$SELECT array_to_string($1, ',')$$;

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

CREATE INDEX "movies_persons_search_idx" ON "movies_persons" USING GIN (
    TO_TSVECTOR(
        'english',
        "movies_persons"."name" || COALESCE(text_array_to_text ("movies_persons"."nicknames"), '') || COALESCE(text_array_to_text ("movies_persons"."occupations"), '') || "movies_persons"."about"
    )
);

-- +goose Down
DROP INDEX "movies_persons_search_idx";

DROP TABLE "movies_persons";

DROP FUNCTION text_array_to_text;
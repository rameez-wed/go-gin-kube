CREATE TABLE "authors" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz DEFAULT NOW(),
  "author_id" bigint NOT NULL
);

CREATE INDEX ON "posts" ("title");

ALTER TABLE "posts" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

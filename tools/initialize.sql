CREATE TABLE "genres" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(255),
  "description" varchar(255)
);

CREATE TABLE "series" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(255),
  "description" varchar(255),
  "genre_id" int references genres(id)
);

CREATE TABLE "seasons" (
  "id" SERIAL PRIMARY KEY,
  "series_id" int references series(id),
  "number" int,
  "name" varchar(255)
);

CREATE TABLE "episodes" (
  "id" SERIAL PRIMARY KEY,
  "season_id" int references seasons(id),
  "name" varchar(255),
  "number" int,
  "description" varchar(255),
  "duration" int,
  "publication_date" date
);

CREATE TABLE "actors" (
  "id" SERIAL PRIMARY KEY,
  "fullname" varchar(255),
  "date_of_birth" date
);

CREATE TABLE "actors_episodes_associations" (
  "actor_id" int references actors(id),
  "episode_id" int references episodes(id)
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "email" varchar(255),
  "full_name" varchar(255),
  "password_hash" varchar(255),
  "password_salt" varchar(255),
  "created_at" timestamp,
  "is_admin" boolean
);

CREATE TABLE "episode_reviews" (
  "id" SERIAL PRIMARY KEY,
  "episode_id" int references episodes(id),
  "reviewer_id" int references users(id),
  "review" text,
  "rate" int,
  "created_at" timestamp
);

CREATE UNIQUE INDEX ON "genres" ("name");

CREATE UNIQUE INDEX ON "series" ("name");

CREATE UNIQUE INDEX ON "users" ("email");


CREATE UNIQUE INDEX "series_season_number" ON "seasons" ("series_id", "number");

CREATE UNIQUE INDEX "season_episode_number" ON "episodes" ("season_id", "number");

CREATE UNIQUE INDEX "actor_episode_presence" ON "actors_episodes_associations" ("actor_id", "episode_id");

CREATE UNIQUE INDEX "user_review" ON "episode_reviews" ("episode_id", "reviewer_id");
CREATE TYPE gender AS ENUM ('male', 'female');
CREATE TABLE "actors" (
  "id" serial,
  "name" varchar NOT NULL,
  "birthday" date NOT NULL,
  "gender" gender NOT NULL
);

CREATE TABLE "films" (
  "id" serial,
  "title" varchar(150) NOT NULL,
  "description" varchar(1000) NOT NULL,
  "release_date" date NOT NULL,
  "rating" smallint NOT NULL
);

CREATE TABLE "actors_films" (
  "actor_id" integer NOT NULL,
  "film_id" integer NOT NULL
);

ALTER TABLE "actors_films" ADD FOREIGN KEY ("actor_id") REFERENCES "actors" ("id");

ALTER TABLE "actors_films" ADD FOREIGN KEY ("film_id") REFERENCES "films" ("id");
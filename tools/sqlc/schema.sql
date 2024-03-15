CREATE TYPE gender AS ENUM ('male', 'female');
CREATE TABLE "actors" (
  "id" serial primary key,
  "name" varchar NOT NULL,
  "birthday" date NOT NULL,
  "gender" gender NOT NULL
);

CREATE TABLE "films" (
  "id" serial primary key,
  "title" varchar(150) NOT NULL,
  "description" varchar(1000) NOT NULL,
  "release_date" date NOT NULL,
  "rating" smallint NOT NULL CHECK(VALUE >= 0 AND VALUE < 100)
);

CREATE TABLE "actors_films" (
  "actor_id" integer NOT NULL,
  "film_id" integer NOT NULL
);

ALTER TABLE "actors_films" ADD FOREIGN KEY ("actor_id") REFERENCES "actors" ("id");

ALTER TABLE "actors_films" ADD FOREIGN KEY ("film_id") REFERENCES "films" ("id");

CREATE TABLE "users" (
    "id" uuid UNIQUE NOT NULL ,
    "nickname" varchar(60) UNIQUE NOT NULL,
    "email" varchar(320) UNIQUE NOT NULL,
    "password_hash" char(64), --SHA-256 hash
    "salt" char(64),
    "is_admin" BOOLEAN
);
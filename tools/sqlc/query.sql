--ACTORS----------------------------------------------------------------------------------------------------------------

-- name: GetActorById :one
SELECT *
FROM actors
WHERE id = $1;

-- name: SearchActorsByName :many
SELECT *
FROM actors
WHERE name LIKE '%' || $1 || '%';

-- name: AddActor :one
INSERT INTO actors (name, birthday, gender)
VALUES ($1, $2, $3)
RETURNING id;

-- name: AddActorToFilm :exec
INSERT INTO actors_films(actor_id, film_id) VALUES ($1, $2);

-- name: updateActorName :exec
UPDATE actors
SET name = $2
WHERE id = $1;

-- name: updateActorBirthday :exec
UPDATE actors
SET birthday = $2
WHERE id = $1;

-- name: updateActorGender :exec
UPDATE actors
SET gender = $2
WHERE id = $1;

-- todo: fix deleting from actors_films
-- name: DeleteActorById :exec
DELETE
FROM actors
WHERE id = $1;


--FILMS-----------------------------------------------------------------------------------------------------------------

-- name: GetFilmById :one
SELECT *
FROM films
WHERE id = $1;

-- name: SearchFilmsByTitle :many
SELECT *
FROM films
WHERE title LIKE '%' || $1 || '%';

-- name: SearchFilmByTitleAndActor :one

WITH ai AS (
	SELECT id FROM actors where lower(name) LIKE lower('%' || $2 || '%')
), fi AS (
	SELECT id FROM films WHERE lower(title) LIKE lower('%' || $1 || '%')
), f AS (
	SELECT film_id FROM actors_films WHERE actor_id IN (select * from ai) AND film_id IN (SELECT * FROM fi)
)
SELECT id, title, description, release_date, rating FROM f JOIN films ON f.film_id = films.id;


-- name: AddFilm :one
INSERT INTO films (title, description, release_date, rating)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: updateFilmTitle :exec
UPDATE films
SET title = $2
WHERE id = $1;

-- name: updateFilmDescription :exec
UPDATE films
SET description = $2
WHERE id = $1;

-- name: updateFilmReleaseDate :exec
UPDATE films
SET release_date = $2
WHERE id = $1;

-- name: updateFilmRating :exec
UPDATE films
SET rating = $2
WHERE id = $1;

-- todo: fix deleting from actors_films
-- name: DeleteFilmById :exec
DELETE
FROM films
WHERE id = $1;

--USERS-----------------------------------------------------------------------------------------------------------------

-- name: GetUserById :one
SELECT * FROM users WHERE id=$1;

-- name: GetUserByLogin :one
SELECT * FROM users WHERE nickname=$1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=$1;

-- name: AddUser :exec
INSERT INTO users (id, nickname, email, password_hash, salt, is_admin) VALUES ($1, $2, $3, $4, $5, $6);

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;
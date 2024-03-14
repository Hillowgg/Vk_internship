--ACTORS----------------------------------------------------------------------------------------------------------------
CREATE FUNCTION search_actors_by_name(arg text)
    returns table
            (
                id integer
            )
AS
$$
SELECT id
FROM actors
WHERE name @@ to_tsquery(arg || ':*')
$$ LANGUAGE SQL;

-- name: GetActorById :one
SELECT *
FROM actors
WHERE id = $1;

-- name SearchActorsByName :many
SELECT *
FROM actors
WHERE name @@ to_tsquery('$1:*');

-- name: AddActor :exec
INSERT INTO actors (name, birthday, gender)
VALUES ($1, $2, $3);

-- name: updateActorName :exec
UPDATE actors
SET name = $2
WHERE id = $1;

-- name: updateActorBirthday :exec
UPDATE actors
SET birdthday = $2
WHERE id = $1;

-- name: updateActorGender :exec
UPDATE actors
SET gender = $2
WHERE id = $1;

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
WHERE title @@ to_tsquery('$1:*');

-- name: SearchFilmsByActorName :many
select *
from films
where films.id IN (
    SELECT film_id FROM actors_films AS t WHERE t.actor_id IN (
        SELECT * FROM search_actors_by_name($1)
    )
);


-- name: AddFilm :exec
INSERT INTO films (title, description, release_date, rating)
VALUES ($1, $2, $3, $3);

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

-- name: DeleteFilmById :exec
DELETE
FROM films
WHERE id = $1;
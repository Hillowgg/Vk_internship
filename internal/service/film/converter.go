package film

import (
    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/database"
)

func newFilmToDbFilm(film *NewFilm) *database.AddFilmParams {
    dbFilm := &database.AddFilmParams{
        Title:       film.Title,
        Description: film.Description,
        ReleaseDate: pgtype.Date{Time: film.ReleaseDate, InfinityModifier: pgtype.Finite, Valid: true},
        Rating:      int16(film.Rating),
    }
    return dbFilm
}

func dbFilmToFilm(dbFilm *database.Film) *Film {
    film := &Film{
        Id:          dbFilm.ID,
        Title:       dbFilm.Title,
        Description: dbFilm.Description,
        ReleaseDate: dbFilm.ReleaseDate.Time,
        Rating:      int8(dbFilm.Rating),
    }
    return film
}

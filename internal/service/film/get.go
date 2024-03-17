package film

import (
    "context"
    "errors"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/database"
    "main/internal/logs"
)

func (s *Service) GetFilm(ctx context.Context, id int32) (*Film, error) {
    dbFilm, err := s.db.GetFilmById(ctx, id)
    if err != nil {
        logs.Log.Errorw("Failed to get a film", "filmId", id, "err", err)
        return nil, err
    }

    logs.Log.Infow("Got film", "filmId", id)
    return dbFilmToFilm(dbFilm), nil
}

func (s *Service) GetFilms(ctx context.Context, sortBy, sortType string) ([]*Film, error) {
    var dbFilms []*database.Film
    var err error
    if sortType == "asc" {
        switch sortBy {
        case "title":
            dbFilms, err = s.db.GetFilmsASCTitle(ctx)
        case "release_date":
            dbFilms, err = s.db.GetFilmsASCReleaseDate(ctx)
        case "rating":
            dbFilms, err = s.db.GetFilmsASCRating(ctx)
        }
    } else {
        switch sortBy {
        case "title":
            dbFilms, err = s.db.GetFilmsDESCTitle(ctx)
        case "release_date":
            dbFilms, err = s.db.GetFilmsDESCReleaseDate(ctx)
        case "rating":
            dbFilms, err = s.db.GetFilmsDESCRating(ctx)
        }
    }

    if err != nil {
        logs.Log.Errorw("Failed to get films", "sortBy", sortBy, "sortType", sortType, "err", err)
        return nil, err
    }
    logs.Log.Infow("Got films", "sortBy", sortBy, "sortType", sortType)
    films := make([]*Film, 0, len(dbFilms))
    for _, film := range dbFilms {
        films = append(films, dbFilmToFilm(film))
    }
    return films, nil
}

func (s *Service) SearchFilms(ctx context.Context, title string) ([]*Film, error) {
    dbFilms, err := s.db.SearchFilmsByTitle(ctx, pgtype.Text{String: title, Valid: true})
    if err != nil {
        logs.Log.Errorw("Failed to search films", "info", title, "err", err)
        return nil, err
    }
    logs.Log.Infow("Found films", "info", title)
    films := make([]*Film, 0, len(dbFilms))
    for _, film := range dbFilms {
        films = append(films, dbFilmToFilm(film))
    }
    return films, nil
}

func (s *Service) SearchFilmByActor(ctx context.Context, title string, actorName string) (*Film, error) {
    dbFilm, err := s.db.SearchFilmByTitleAndActor(
        ctx,
        &database.SearchFilmByTitleAndActorParams{
            Column1: pgtype.Text{String: title, Valid: true},
            Column2: pgtype.Text{String: actorName, Valid: true},
        })
    logs.Log.Infow("Searching for film with title and actor's name",
        "title", title, "name", actorName)
    if errors.Is(err, pgx.ErrNoRows) {
        return nil, nil
    }
    if err != nil {
        logs.Log.Errorw("Failed to Search film with title and actor",
            "title", title, "name", actorName, "err", err)
        return nil, err
    }

    return dbFilmToFilm(dbFilm), nil
}

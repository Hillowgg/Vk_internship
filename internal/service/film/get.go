package film

import (
    "context"

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
    if err != nil {
        logs.Log.Errorw("Failed to Search film with title and actor",
            "title", title, "name", actorName, "err", err)
        return nil, err
    }
    logs.Log.Infow("Searched for film with title and actor's name",
        "title", title, "name", actorName)
    return dbFilmToFilm(dbFilm), nil
}

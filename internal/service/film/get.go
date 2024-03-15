package film

import (
    "context"

    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/database"
    "main/internal/logs"
)

func (s *Service) GetFilm(ctx context.Context, id int32) (*database.Film, error) {
    film, err := s.db.GetFilmById(ctx, id)
    if err != nil {
        logs.Log.Errorw("Failed to get a film", "filmId", id, "err", err)
        return nil, err
    }
    logs.Log.Infow("Got film", "filmId", id)
    return film, nil
}

func (s *Service) SearchFilms(ctx context.Context, title string) ([]*database.Film, error) {
    films, err := s.db.SearchFilmsByTitle(ctx, pgtype.Text{String: title, Valid: true})
    if err != nil {
        logs.Log.Errorw("Failed to search films", "info", title, "err", err)
        return nil, err
    }
    logs.Log.Infow("Found films", "info", title)
    return films, nil
}

func (s *Service) SearchFilmByActor(ctx context.Context, title string, actorName string) (*database.Film, error) {
    film, err := s.db.SearchFilmByTitleAndActor(
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
    return film, nil
}

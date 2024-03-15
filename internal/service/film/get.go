package film

import (
    "context"

    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/database"
)

// GetFilm todo: reallocation on heap
func (s *Service) GetFilm(ctx context.Context, id int32) (database.Film, error) {
    film, err := s.db.GetFilmById(ctx, id)
    if err != nil {
        return database.Film{}, err
    }
    return film, err
}

func (s *Service) SearchFilms(ctx context.Context, title string) ([]database.Film, error) {
    films, err := s.db.SearchFilmsByTitle(ctx, pgtype.Text{String: title, Valid: true})
    if err != nil {
        return nil, err
    }
    return films, nil
}

func (s *Service) SearchFilmByActor(ctx context.Context, title string, actorName string) (database.Film, error) {
    film, err := s.db.SearchFilmByTitleAndActor(
        ctx,
        database.SearchFilmByTitleAndActorParams{
            Column1: pgtype.Text{String: title, Valid: true},
            Column2: pgtype.Text{String: actorName, Valid: true},
        })
    if err != nil {
        return database.Film{}, err
    }
    return film, nil
}

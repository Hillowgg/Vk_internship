package film

import (
    "context"

    "main/internal/database"
    "main/internal/logs"
)

type IService interface {
    GetFilm(ctx context.Context, id int32) (database.Film, error)
    SearchFilms(ctx context.Context, title string) ([]database.Film, error)
    SearchFilmByActor(ctx context.Context, title string, actorName string) (database.Film, error)

    CreateFilm(ctx context.Context, film database.AddFilmParams) (int32, error)
    CreateFilmWithActors(ctx context.Context, film database.AddFilmParams, actors []int32) error
    UpdateFilm(ctx context.Context, id int32, newFilm database.OptUpdateFilm) error

    DeleteFilm(ctx context.Context, id int32) error
}

type Service struct {
    db *database.Queries
}

func NewService(db *database.Queries) (*Service, error) {
    if db == nil {
        logs.Log.Error("Failed to create Film service: db is nil")
    }
    return &Service{db}, nil
}

package film

import (
    "context"
    "time"

    "main/internal/database"
    "main/internal/logs"
)

type Film struct {
    Id          int32
    Title       string
    Description string
    ReleaseDate time.Time
    Rating      int8
}

type NewFilm struct {
    Title       string
    Description string
    ReleaseDate time.Time
    Rating      int8
}
type IService interface {
    GetFilm(ctx context.Context, id int32) (*Film, error)
    GetFilms(ctx context.Context, sortBy, sortType string) ([]*Film, error)
    SearchFilms(ctx context.Context, title string) ([]*Film, error)
    SearchFilmByActor(ctx context.Context, title string, actorName string) (*Film, error)

    CreateFilm(ctx context.Context, film *NewFilm) (int32, error)
    CreateFilmWithActors(ctx context.Context, film *NewFilm, actors []int32) (int32, error)
    UpdateFilm(ctx context.Context, film map[string]any) error

    DeleteFilm(ctx context.Context, id int32) error
}

type Service struct {
    db database.FilmQuerierWithTx
}

func NewService(db database.FilmQuerierWithTx) *Service {
    if db == nil {
        logs.Log.Error("Failed to create Film service: db is nil")
    }
    return &Service{db}
}

var _ IService = (*Service)(nil)

package actor

import (
    "context"
    "time"

    "main/internal/database"
    "main/internal/logs"
    "main/internal/service/film"
)

type Actor struct {
    Id       int32
    Name     string
    Birthday time.Time
    Gender   database.Gender
}

type NewActor struct {
    Name     string
    Birthday time.Time
    Gender   database.Gender
}

type IService interface {
    GetActor(ctx context.Context, id int32) (*Actor, error)
    GetActorsWithFilms(ctx context.Context) (map[Actor][]film.Film, error)
    SearchActors(ctx context.Context, name string) ([]*Actor, error)

    CreateActor(ctx context.Context, actor *NewActor) (int32, error)
    AddActorToFilm(ctx context.Context, actorId, filmId int32) error

    UpdateActor(ctx context.Context, actor map[string]any) error

    DeleteActor(ctx context.Context, id int32) error
}

type Service struct {
    db database.ActorQuerierWithTx
}

func NewService(db database.ActorQuerierWithTx) *Service {
    if db == nil {
        logs.Log.Fatal("Failed to create service: db is nil")
    }
    return &Service{db: db}
}

var _ IService = (*Service)(nil)

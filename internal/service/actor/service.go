package actor

import (
    "context"
    "time"

    "main/internal/database"
    "main/internal/logs"
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
    SearchActors(ctx context.Context, name string) ([]*Actor, error)

    CreateActor(ctx context.Context, actor *NewActor) (int32, error)
    AddActorToFilm(ctx context.Context, actorId, filmId int32) error

    UpdateActor(ctx context.Context, actor *database.OptUpdateActor) error

    DeleteActor(ctx context.Context, id int32) error
}

type Service struct {
    db database.QuerierWithTx
}

func NewService(db database.QuerierWithTx) *Service {
    if db == nil {
        logs.Log.Fatal("Failed to create service: db is nil")
    }
    return &Service{db: db}
}

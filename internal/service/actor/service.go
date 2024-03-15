package actor

import (
    "context"

    "main/internal/database"
)

type IService interface {
    GetActor(ctx context.Context, id int32) (database.Actor, error)
    SearchActors(ctx context.Context, id int32) ([]database.Actor, error)

    CreateActor(ctx context.Context, params database.AddActorParams) (int32, error)
    AddActorToFilm(ctx context.Context, actorId, filmId int32) error

    UpdateActor(ctx context.Context, id int32, film database.OptUpdateActor) error

    DeleteActor(ctx context.Context) error
}

type Service struct {
    db *database.Queries
}

func NewService(db *database.Queries) *Service {
    return &Service{db: db}
}

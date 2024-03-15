package actor

import (
    "context"

    "main/internal/database"
)

func (s *Service) AddActorToFilm(ctx context.Context, actorId, filmId int32) error {
    // TODO implement me
    panic("implement me")
}

func (s *Service) UpdateActor(ctx context.Context, id int32, film database.OptUpdateActor) error {
    // TODO implement me
    panic("implement me")
}

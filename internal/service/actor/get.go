package actor

import (
    "context"

    "main/internal/database"
)

func (s *Service) GetActor(ctx context.Context, id int32) (database.Actor, error) {
    // TODO implement me
    panic("implement me")
}

func (s *Service) SearchActors(ctx context.Context, id int32) ([]database.Actor, error) {
    // TODO implement me
    panic("implement me")
}

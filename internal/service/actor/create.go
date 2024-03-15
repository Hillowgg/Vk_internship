package actor

import (
    "context"

    "main/internal/database"
)

func (s *Service) CreateActor(ctx context.Context, actor database.AddActorParams) (int32, error) {
    id, err := s.db.AddActor(ctx, actor)
    return id, err
}

package actor

import (
    "context"

    "main/internal/database"
    "main/internal/logs"
)

func (s *Service) CreateActor(ctx context.Context, actor *database.AddActorParams) (int32, error) {
    id, err := s.db.AddActor(ctx, actor)
    if err != nil {
        logs.Log.Errorw("Failed to create an actor", "info", actor, "err", err)
        return 0, err
    }
    logs.Log.Infow("Created an actor", "actorId", id)
    return id, nil
}

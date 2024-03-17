package actor

import (
    "context"
    "strings"

    "main/internal/logs"
)

func (s *Service) CreateActor(ctx context.Context, actor *NewActor) (int32, error) {
    actor.Name = strings.ToLower(actor.Name)
    id, err := s.db.AddActor(ctx, newActorToDb(actor))
    if err != nil {
        logs.Log.Errorw("Failed to create an actor", "info", actor, "err", err)
        return 0, err
    }
    logs.Log.Infow("Created an actor", "actorId", id)
    return id, nil
}

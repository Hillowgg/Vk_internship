package actor

import (
    "context"

    "main/internal/logs"
)

func (s *Service) DeleteActor(ctx context.Context, id int32) error {
    err := s.db.DeleteActorById(ctx, id)
    if err != nil {
        logs.Log.Errorw("Failed to delete an actor", "actorId", id, "err", err)
        return err
    }
    logs.Log.Infow("Deleted actor", "actorId", id)
    return nil
}

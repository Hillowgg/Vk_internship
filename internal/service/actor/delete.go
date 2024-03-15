package actor

import (
    "context"
)

func (s *Service) DeleteActor(ctx context.Context, id int32) error {
    err := s.db.DeleteActorById(ctx, id)
    return err
}

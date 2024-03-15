package actor

import (
    "context"

    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/database"
    "main/internal/logs"
)

func (s *Service) GetActor(ctx context.Context, id int32) (*database.Actor, error) {
    actor, err := s.db.GetActorById(ctx, id)
    if err != nil {
        logs.Log.Errorw("Failed to get an actor", "actorId", id, "err", err)
        return nil, err
    }
    logs.Log.Infow("Got an actor", "actorId", id)
    return actor, nil
}

func (s *Service) SearchActors(ctx context.Context, name string) ([]*database.Actor, error) {
    actors, err := s.db.SearchActorsByName(ctx, pgtype.Text{name, true})
    if err != nil {
        logs.Log.Errorw("Failed to search actors", "info", name, "err", err)
        return nil, err
    }
    logs.Log.Infow("Searched actors", "info", name)
    return actors, nil
}

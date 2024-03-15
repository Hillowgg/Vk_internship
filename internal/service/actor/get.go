package actor

import (
    "context"

    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/logs"
)

func (s *Service) GetActor(ctx context.Context, id int32) (*Actor, error) {
    actor, err := s.db.GetActorById(ctx, id)
    if err != nil {
        logs.Log.Errorw("Failed to get an actor", "actorId", id, "err", err)
        return nil, err
    }
    logs.Log.Infow("Got an actor", "actorId", id)
    return dbActorToActor(actor), nil
}

func (s *Service) SearchActors(ctx context.Context, name string) ([]*Actor, error) {
    actors, err := s.db.SearchActorsByName(ctx, pgtype.Text{name, true})
    if err != nil {
        logs.Log.Errorw("Failed to search actors", "info", name, "err", err)
        return nil, err
    }
    logs.Log.Infow("Searched actors", "info", name)
    ret := make([]*Actor, 0, len(actors))
    for _, actor := range actors {
        ret = append(ret, dbActorToActor(actor))
    }
    return ret, nil
}

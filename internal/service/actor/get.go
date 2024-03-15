package actor

import (
    "context"

    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/database"
)

func (s *Service) GetActor(ctx context.Context, id int32) (database.Actor, error) {
    actor, err := s.db.GetActorById(ctx, id)
    return actor, err
}

func (s *Service) SearchActors(ctx context.Context, name string) ([]database.Actor, error) {
    actors, err := s.db.SearchActorsByName(ctx, pgtype.Text{name, true})
    return actors, err
}

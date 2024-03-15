package actor

import (
    "context"

    "main/internal/database"
    "main/internal/logs"
)

func (s *Service) AddActorToFilm(ctx context.Context, actorId, filmId int32) error {
    err := s.db.AddActorToFilm(ctx, database.AddActorToFilmParams{ActorID: actorId, FilmID: filmId})
    if err != nil {
        logs.Log.Errorw(
            "Failed to add actor to film",
            "actorId", actorId,
            "filmId", filmId,
            "err", err,
        )
        return err
    }
    logs.Log.Infow("Added actor to film", "actorId", actorId, "filmId", filmId)
    return nil
}

func (s *Service) UpdateActor(ctx context.Context, id int32, actor database.OptUpdateActor) error {
    err := s.db.UpdateActor(ctx, actor)
    if err != nil {
        logs.Log.Errorw("Failed to update actor", "actorId", id, "info", actor, "err", err)
        return err
    }
    logs.Log.Infow("Updated actor", "actorId", id, "info", actor)
    return nil
}

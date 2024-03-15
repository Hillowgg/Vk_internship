package actor

import (
    "context"

    "main/internal/database"
)

func (s *Service) AddActorToFilm(ctx context.Context, actorId, filmId int32) error {
    err := s.db.AddActorToFilm(ctx, database.AddActorToFilmParams{ActorID: actorId, FilmID: filmId})
    return err
}

func (s *Service) UpdateActor(ctx context.Context, id int32, actor database.OptUpdateActor) error {
    err := s.db.UpdateActor(ctx, actor)
    return err
}

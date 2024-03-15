package film

import (
    "context"

    "main/internal/logs"
)

func (s *Service) CreateFilm(ctx context.Context, film *NewFilm) (int32, error) {

    id, err := s.db.AddFilm(ctx, newFilmToDbFilm(film))
    if err != nil {
        logs.Log.Errorw("Failed to create film", "info", film, "err", err)
        return 0, err
    }
    logs.Log.Infow("Added film", "filmId", id)
    return id, nil
}

// TODO: add returning id
func (s *Service) CreateFilmWithActors(ctx context.Context, film *NewFilm, actors []int32) error {
    err := s.db.AddFilmWithActors(ctx, newFilmToDbFilm(film), actors)
    if err != nil {
        logs.Log.Errorw("Failed to create film with actors",
            "filmInfo", film, "actorIds", actors, "err", err)
        return err
    }
    logs.Log.Infow("Created film with actors")
    return nil
}

package film

import (
    "context"

    "main/internal/database"
    "main/internal/logs"
)

func (s *Service) CreateFilm(ctx context.Context, film *database.AddFilmParams) (int32, error) {
    id, err := s.db.AddFilm(ctx, film)
    if err != nil {
        logs.Log.Errorw("Failed to create film", "info", film, "err", err)
        return 0, err
    }
    logs.Log.Infow("Added film", "filmId", id)
    return id, nil
}

// TODO: add returning id
func (s *Service) CreateFilmWithActors(ctx context.Context, film database.AddFilmParams, actors []int32) error {
    err := s.db.AddFilmWithActors(ctx, film, actors)
    if err != nil {
        logs.Log.Errorw("Failed to create film with actors",
            "filmInfo", film, "actorIds", actors, "err", err)
        return err
    }
    logs.Log.Infow("Created film with actors")
    return nil
}

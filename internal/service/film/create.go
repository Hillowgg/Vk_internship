package film

import (
    "context"

    "main/internal/database"
)

func (s *Service) CreateFilm(ctx context.Context, film database.AddFilmParams) (int32, error) {
    id, err := s.db.AddFilm(ctx, film)
    return id, err
}

func (s *Service) CreateFilmWithActors(ctx context.Context, film database.AddFilmParams, actors []int32) error {
    err := s.db.AddFilmWithActors(ctx, film, actors)
    return err
}

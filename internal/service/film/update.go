package film

import (
    "context"

    "main/internal/database"
    "main/internal/logs"
)

// UpdateFilm nil stands for 'not to update'
func (s *Service) UpdateFilm(ctx context.Context, id int32, film database.OptUpdateFilm) error {
    err := s.db.UpdateFilm(ctx, film)
    if err != nil {
        logs.Log.Errorw("Failed to update film", "filmId", id, "err", err)
        return err
    }
    logs.Log.Infow("Updated film", "filmId", id)
    return err
}

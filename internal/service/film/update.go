package film

import (
    "context"

    "main/internal/database"
)

// UpdateFilm nil stands for 'not to update'
func (s *Service) UpdateFilm(ctx context.Context, id int32, film database.OptUpdateFilm) error {
    err := s.db.UpdateFilm(ctx, film)
    return err
}

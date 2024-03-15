package film

import (
    "context"

    "main/internal/logs"
)

func (s *Service) DeleteFilm(ctx context.Context, id int32) error {
    err := s.db.DeleteFilmById(ctx, id)
    if err != nil {
        logs.Log.Errorw("Failed to delete film", "filmId", id, "err", err)
        return err
    }
    logs.Log.Infow("Deleted film", "filmId", id)
    return nil
}

package film

import (
    "context"
)

func (s *Service) DeleteFilm(ctx context.Context, id int32) error {
    err := s.db.DeleteFilmById(ctx, id)
    return err
}

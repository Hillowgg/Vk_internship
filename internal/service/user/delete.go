package user

import (
    "context"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/logs"
)

func (s Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
    bytes := [16]byte(id)
    err := s.db.DeleteUser(ctx, pgtype.UUID{Bytes: bytes, Valid: true})
    if err != nil {
        logs.Log.Errorw("Failed to delete user", "uuid", id, "err", err)
        return err
    }
    logs.Log.Infow("Deleted user", "uuid", id)
    return nil
}

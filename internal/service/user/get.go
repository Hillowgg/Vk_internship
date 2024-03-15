package user

import (
    "context"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/database"
    "main/internal/logs"
)

func (s Service) GetUser(ctx context.Context, id uuid.UUID) (*database.User, error) {
    bytes := [16]byte(id)
    user, err := s.db.GetUserById(ctx, pgtype.UUID{Bytes: bytes, Valid: true})
    if err != nil {
        logs.Log.Errorw("Failed to get user", "uuid", id, "err", err)
        return nil, err
    }
    return user, nil
}

func (s Service) IsUserAdmin(ctx context.Context, id uuid.UUID) (bool, error) {
    user, err := s.GetUser(ctx, id)
    v := user.IsAdmin
    return v.Bool && v.Valid, err
}

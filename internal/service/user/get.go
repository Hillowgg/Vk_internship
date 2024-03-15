package user

import (
    "context"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/logs"
)

func (s Service) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
    bytes := [16]byte(id)
    user, err := s.db.GetUserById(ctx, pgtype.UUID{Bytes: bytes, Valid: true})
    if err != nil {
        logs.Log.Errorw("Failed to get user", "uuid", id, "err", err)
        return nil, err
    }
    id, err = uuid.FromBytes(user.ID.Bytes[:])
    if err != nil {
        logs.Log.Errorw("Failed to parse uuid while getting user", "err", err)
        return nil, err
    }

    res := &User{
        Id:       id,
        Nickname: user.Nickname,
        Email:    user.Email,
        Password: "",
        IsAdmin:  user.IsAdmin.Bool,
    }
    return res, nil
}

func (s Service) IsUserAdmin(ctx context.Context, id uuid.UUID) (bool, error) {
    user, err := s.GetUser(ctx, id)

    return user.IsAdmin, err
}

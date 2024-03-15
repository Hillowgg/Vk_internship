package user

import (
    "context"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/database"
    "main/internal/logs"
)

func (s Service) CreateUser(ctx context.Context, user *NewUser) (uuid.UUID, error) {
    id := uuid.New()
    hash, salt, err := generateHashedPassword(user.Password)
    if err != nil {
        logs.Log.Errorw("Error while generating salt", "err", err)
        return [16]byte{}, err
    }

    u := &database.AddUserParams{
        ID:           pgtype.UUID{Bytes: [16]byte(id), Valid: true},
        Nickname:     user.Nickname,
        Email:        user.Email,
        IsAdmin:      pgtype.Bool{Bool: user.IsAdmin, Valid: true},
        PasswordHash: pgtype.Text{String: hash, Valid: true},
        Salt:         pgtype.Text{String: salt, Valid: true},
    }

    err = s.db.AddUser(ctx, u)
    if err != nil {
        logs.Log.Errorw("Failed to create user", "info", user, "err", err)
        return [16]byte{}, err
    }
    logs.Log.Infow("Created user", "info", user)
    return id, err
}

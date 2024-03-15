package user

import (
    "context"

    "main/internal/database"
    "main/internal/logs"
)

func (s Service) CreateUser(ctx context.Context, user *database.AddUserParams) error {
    err := s.db.AddUser(ctx, user)
    if err != nil {
        logs.Log.Errorw("Failed to create user", "info", user, "err", err)
        return err
    }
    logs.Log.Infow("Created user", "info", user)
    return nil
}

package user

import (
    "context"

    "github.com/google/uuid"
    "main/internal/database"
    "main/internal/logs"
)

type IService interface {
    GetUser(ctx context.Context, id uuid.UUID) (*database.User, error)
    IsUserAdmin(ctx context.Context, id uuid.UUID) (bool, error)

    CreateUser(ctx context.Context, user *database.AddUserParams) error

    DeleteUser(ctx context.Context, id uuid.UUID) error
}

type Service struct {
    db database.Querier
}

func NewService(db database.Querier) *Service {
    if db == nil {
        logs.Log.Fatal("Failed to create user serivce: db is nil")
    }
    return &Service{db: db}
}

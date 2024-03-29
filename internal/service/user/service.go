package user

import (
    "context"

    "github.com/google/uuid"
    "main/internal/database"
    "main/internal/logs"
)

type User struct {
    Id       uuid.UUID
    Nickname string
    Email    string
    Password string
    IsAdmin  bool
}

type NewUser struct {
    Nickname string
    Email    string
    Password string
    IsAdmin  bool
}

type IService interface {
    GetUser(ctx context.Context, id uuid.UUID) (*User, error)
    IsUserAdmin(ctx context.Context, id uuid.UUID) (bool, error)
    CheckEmailCredentials(ctx context.Context, email, password string) (*User, error)
    CheckLoginCredentials(ctx context.Context, login, password string) (*User, error)

    CreateUser(ctx context.Context, user *NewUser) (uuid.UUID, error)

    DeleteUser(ctx context.Context, id uuid.UUID) error
}

type Service struct {
    db database.UserQuerierWithTx
}

func NewService(db database.UserQuerierWithTx) *Service {
    if db == nil {
        logs.Log.Fatal("Failed to create user serivce: db is nil")
    }
    return &Service{db: db}
}

package user

import (
    "context"
    "errors"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgtype"
    "main/internal/database"
    "main/internal/logs"
)

func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
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

func (s *Service) IsUserAdmin(ctx context.Context, id uuid.UUID) (bool, error) {
    user, err := s.GetUser(ctx, id)

    return user.IsAdmin, err
}

var WrongCredentials = errors.New("wrong credentials")

func checkCredentials(dbUser *database.User, password string) (*User, error) {
    if !checkPassword(password, dbUser.Salt.String, dbUser.PasswordHash.String) {
        return nil, WrongCredentials
    }
    logs.Log.Infow("Checked credentials", "user", dbUser, "password", password)
    return dbUserToUser(dbUser), nil
}
func (s *Service) CheckEmailCredentials(ctx context.Context, email, password string) (*User, error) {
    dbUser, err := s.db.GetUserByEmail(ctx, email)
    if errors.Is(err, pgx.ErrNoRows) {
        return nil, WrongCredentials
    } else if err != nil {
        logs.Log.Errorw("Failed to check email credentials",
            "email", email, "password", password, "err", err)
        return nil, err
    }
    return checkCredentials(dbUser, password)
}
func (s *Service) CheckLoginCredentials(ctx context.Context, login, password string) (*User, error) {
    dbUser, err := s.db.GetUserByLogin(ctx, login)
    if errors.Is(err, pgx.ErrNoRows) {
        return nil, WrongCredentials
    } else if err != nil {
        logs.Log.Errorw("Failed to check login credentials",
            "login", login, "password", password, "err", err)
        return nil, err
    }
    return checkCredentials(dbUser, password)
}

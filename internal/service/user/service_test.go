package user

import (
    "context"
    "errors"
    "testing"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgtype"
    "github.com/stretchr/testify/mock"
    "main/internal/database"
)

// write tests for user service

type mockDB struct {
    mock.Mock
}

func (m *mockDB) AddUser(ctx context.Context, user *database.AddUserParams) error {
    args := m.Called(ctx, user)
    return args.Error(0)
}
func (m *mockDB) GetUserByEmail(ctx context.Context, email string) (*database.User, error) {
    args := m.Called(ctx, email)
    return args.Get(0).(*database.User), args.Error(1)
}
func (m *mockDB) GetUserByLogin(ctx context.Context, nickname string) (*database.User, error) {
    args := m.Called(ctx, nickname)
    return args.Get(0).(*database.User), args.Error(1)
}

func (m *mockDB) GetUserById(ctx context.Context, id pgtype.UUID) (*database.User, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*database.User), args.Error(1)
}

func (m *mockDB) DeleteUser(ctx context.Context, id pgtype.UUID) error {
    args := m.Called(ctx, id)
    return args.Error(0)
}

func TestService_GetUser(t *testing.T) {
    // write test for GetUser
    db := new(mockDB)
    s := NewService(db)
    u := &User{}
    db.On("GetUserById", mock.Anything, mock.Anything).Return(u, nil)
    _, err := s.GetUser(context.Background(), uuid.New())
    if err != nil {
        t.Error("GetUser returned error")
    }
}

func TestService_CreateUser(t *testing.T) {
    // write test for CreateUser
    db := new(mockDB)
    s := NewService(db)
    db.On("AddUser", mock.Anything, mock.Anything).Return(nil)
    _, err := s.CreateUser(context.Background(), &NewUser{})
    if err != nil {
        t.Error("CreateUser returned error")
    }
}

func TestService_DeleteUser(t *testing.T) {
    // write test for DeleteUser
    db := new(mockDB)
    s := NewService(db)
    db.On("DeleteUser", mock.Anything, mock.Anything).Return(nil)
    err := s.DeleteUser(context.Background(), uuid.New())
    if err != nil {
        t.Error("DeleteUser returned error")
    }
}

func TestService_CheckEmailCredentials(t *testing.T) {
    // write test for CheckEmailCredentials
    db := new(mockDB)
    s := NewService(db)
    db.On("GetUserByEmail", mock.Anything, mock.Anything).Return(&database.User{}, nil)
    _, err := s.CheckEmailCredentials(context.Background(), "email", "password")
    if err != nil && !errors.Is(err, WrongCredentials) {
        t.Error("CheckEmailCredentials returned error")
    }
}

func TestService_CheckLoginCredentials(t *testing.T) {
    // write test for CheckLoginCredentials
    db := new(mockDB)
    s := NewService(db)
    db.On("GetUserByLogin", mock.Anything, mock.Anything).Return(&database.User{}, nil)
    _, err := s.CheckLoginCredentials(context.Background(), "login", "password")
    if err != nil && !errors.Is(err, WrongCredentials) {
        t.Error("CheckLoginCredentials returned error")
    }
}

func TestService_IsUserAdmin(t *testing.T) {
    // write test for IsUserAdmin
    db := new(mockDB)
    s := NewService(db)
    db.On("GetUserById", mock.Anything, mock.Anything).Return(&database.User{IsAdmin: pgtype.Bool{Bool: true, Valid: true}}, nil)
    admin, err := s.IsUserAdmin(context.Background(), uuid.New())
    if err != nil {
        t.Error("IsUserAdmin returned error")
    }
    if !admin {
        t.Error("IsUserAdmin returned false")
    }
}

// go test -v ./internal/service/user

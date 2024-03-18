package session

import (
    "context"
    "testing"

    "github.com/google/uuid"
    "github.com/redis/go-redis/v9"
    "github.com/stretchr/testify/mock"
    "main/internal/cache"
)

type mockCache struct {
    mock.Mock
}

var NIL = (*mockCache)(nil)

func (m *mockCache) NewSession(ctx context.Context, token string, session cache.Session) error {
    args := m.Called(ctx, token, session)
    return args.Error(0)
}

func (m *mockCache) GetSession(ctx context.Context, token string) (*cache.Session, error) {
    args := m.Called(ctx, token)
    if args.Error(1) != nil {
        return nil, args.Error(1)
    }

    return args.Get(0).(*cache.Session), args.Error(1)
}

func TestService_CreateSession(t *testing.T) {
    // write test for CreateSession
    c := new(mockCache)
    s := NewService(c)
    c.On("NewSession", mock.Anything, mock.Anything, mock.Anything).Return(nil)
    _, err := s.CreateSession(context.Background(), uuid.New(), false)
    if err != nil {
        t.Error("CreateSession returned error")
    }
}

func TestService_GetSession(t *testing.T) {
    // write test for GetSession
    c := new(mockCache)
    if c == NIL {
        return
    }
    s := NewService(c)
    c.On("GetSession", mock.Anything, mock.Anything).Return(&cache.Session{}, nil)
    _, err := s.GetSession(context.Background(), "token")
    if err != nil {
        t.Error("GetSession returned error")
    }
}

func TestService_IsAdmin(t *testing.T) {
    // write test for IsAdmin
    c := new(mockCache)
    if c == NIL {
        return
    }
    s := NewService(c)
    c.On("GetSession", mock.Anything, mock.Anything).Return(&cache.Session{IsAdmin: true}, nil)
    if !s.IsAdmin(context.Background(), "token") {
        t.Error("IsAdmin returned false")
    }
}

func TestService_IsUser(t *testing.T) {
    // write test for IsUser
    c := new(mockCache)
    if c == NIL {
        return
    }
    s := NewService(c)
    c.On("GetSession", mock.Anything, mock.Anything).Return(&cache.Session{IsAdmin: false}, nil)
    if !s.IsUser(context.Background(), "token") {
        t.Error("IsUser returned false")
    }
}

func TestService_IsAdminFalse(t *testing.T) {
    // write test for IsAdmin
    c := new(mockCache)
    if c == NIL {
        return
    }
    s := NewService(c)
    c.On("GetSession", mock.Anything, mock.Anything).Return(&cache.Session{IsAdmin: false}, nil)
    if s.IsAdmin(context.Background(), "token") {
        t.Error("IsAdmin returned true")
    }
}

func TestService_GetSessionError(t *testing.T) {
    // write test for GetSession
    c := new(mockCache)
    s := NewService(c)
    if c == NIL {
        return
    }
    c.On("GetSession", mock.Anything, mock.Anything).Return(nil, redis.Nil)
    _, err := s.GetSession(context.Background(), "token")
    if err != nil {
        t.Error("GetSession returned nil error")
    }
}

func TestService_CreateSessionError(t *testing.T) {

    c := new(mockCache)
    s := NewService(c)
    if c == NIL {
        return
    }
    c.On("NewSession", mock.Anything, mock.Anything, mock.Anything).Return(redis.Nil)
    _, err := s.CreateSession(context.Background(), uuid.New(), false)

    if err == nil {
        t.Error("CreateSession returned nil error")
    }
}

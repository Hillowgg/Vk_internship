package session

import (
    "context"

    "github.com/google/uuid"
    "main/internal/cache"
)

// IService is an interface for session service
type IService interface {
    CreateSession(ctx context.Context, id uuid.UUID, isAdmin bool) (string, error)
    GetSession(ctx context.Context, token string) (*cache.Session, error)
    IsAdmin(ctx context.Context, token string) bool
    IsUser(ctx context.Context, token string) bool
}

// Service is a session service
type Service struct {
    cache cache.ICache
}

// NewService creates a new session service
func NewService(cache cache.ICache) *Service {
    return &Service{cache: cache}
}

var _ IService = &Service{}

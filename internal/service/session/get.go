package session

import (
    "context"
    "errors"

    "github.com/redis/go-redis/v9"
    "main/internal/cache"
    "main/internal/logs"
)

func (s *Service) IsAdmin(ctx context.Context, token string) bool {
    session, err := s.GetSession(ctx, token)
    logs.Log.Infow("checking session for admin", "session", session)
    if err != nil {
        logs.Log.Errorw("failed to get session", "err", err)
        return false
    }
    return session.IsAdmin
}
func (s *Service) IsUser(ctx context.Context, token string) bool {
    _, err := s.GetSession(ctx, token)
    logs.Log.Infow("checked session for user", "token", token)
    if err != nil {
        logs.Log.Errorw("failed to get session", "err", err)
        return false
    }
    return true
}

func (s *Service) GetSession(ctx context.Context, token string) (*cache.Session, error) {
    logs.Log.Infow("getting session", "token", token)
    ses, err := s.cache.GetSession(ctx, token)
    if errors.Is(err, redis.Nil) {
        logs.Log.Infow("session not found", "token", token)
        return nil, nil
    }
    if err != nil {
        logs.Log.Errorw("failed to get session", "token", token, "err", err)
        return nil, err
    }
    return ses, nil
}

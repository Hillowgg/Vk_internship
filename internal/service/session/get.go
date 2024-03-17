package session

import (
    "context"

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
    logs.Log.Infow("checked session for user", "err", err)
    if err != nil {
        logs.Log.Errorw("failed to get session", "err", err)
        return false
    }
    return true
}

func (s *Service) GetSession(ctx context.Context, token string) (*cache.Session, error) {
    logs.Log.Infow("getting session", "token", token)
    return s.cache.GetSession(ctx, token)
}

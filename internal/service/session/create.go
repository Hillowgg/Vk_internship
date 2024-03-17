package session

import (
    "context"
    "crypto/rand"
    "encoding/hex"

    "github.com/google/uuid"
    "main/internal/cache"
    "main/internal/logs"
)

func randomHex(n int) (string, error) {
    bytes := make([]byte, n)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

func (s *Service) CreateSession(ctx context.Context, id uuid.UUID, isAdmin bool) (string, error) {
    token, _ := randomHex(20)
    session := cache.Session{UserID: id, IsAdmin: isAdmin}
    err := s.cache.NewSession(ctx, token, session)
    if err != nil {
        logs.Log.Errorw("failed to create session", "err", err)
        return "", err
    }
    logs.Log.Infow("created session", "token", token)
    return token, nil
}

package cache

import (
    "context"
    "encoding/json"
    "errors"
    "time"

    "github.com/google/uuid"
)

type Session struct {
    UserID  uuid.UUID
    IsAdmin bool
}

func (s Session) MarshalBinary() ([]byte, error) {
    return json.Marshal(s)
}

func (c *Cache) GetSession(ctx context.Context, token string) (*Session, error) {
    res := c.redis.Get(ctx, "session:"+token)
    if res.Err() != nil {
        return nil, errors.New("session not found")
    }
    var s Session
    err := json.Unmarshal([]byte(res.Val()), &s)
    return &s, err
}

func (c *Cache) NewSession(ctx context.Context, token string, session Session) error {
    return c.redis.Set(ctx, "session:"+token, session, time.Hour*24).Err()
}

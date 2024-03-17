package cache

import (
    "context"

    redis2 "github.com/redis/go-redis/v9"
    "main/internal/logs"
)

type ICache interface {
    GetSession(ctx context.Context, token string) (*Session, error)
    NewSession(ctx context.Context, token string, session Session) error
}

type Cache struct {
    redis *redis2.Client
}

func NewCache() *Cache {
    opts, err := redis2.ParseURL("redis://127.0.0.1:6379")
    redis := redis2.NewClient(opts)
    if err != nil {
        logs.Log.Fatalw("Failed to connect to redis", "err", err)
    }
    return &Cache{redis: redis}
}

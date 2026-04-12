package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
}

type cacheImpl struct {
	// redis
	redis *redis.Client
}

func NewCache(r *redis.Client) Cache {
	return &cacheImpl{redis: r}
}

func (c *cacheImpl) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	jsonData, err := json.Marshal(value)

	if err != nil {
		return err
	}

    return c.redis.Set(ctx, key, jsonData, expiration).Err()
}

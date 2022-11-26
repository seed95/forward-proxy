package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/seed95/forward-proxy/internal/repo"
	"time"
)

type redisRepo struct {
	rdb               *redis.Client
	defaultExpiration time.Duration
}

var _ repo.CacheRepo = (*redisRepo)(nil)

func New(rdb *redis.Client, expiration time.Duration) repo.CacheRepo {
	return &redisRepo{
		rdb:               rdb,
		defaultExpiration: expiration,
	}
}

func (r *redisRepo) CacheResponse(ctx context.Context, url string, body []byte) (err error) {
	return r.rdb.Set(ctx, url, string(body), r.defaultExpiration).Err()
}

func (r *redisRepo) GetCachedRequest(ctx context.Context, url string) (body []byte) {
	result, err := r.rdb.Get(ctx, url).Result()
	if err == redis.Nil {
		return nil
	}
	fmt.Println("redis", result)
	return []byte(result)
}

package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/seed95/forward-proxy/internal"
	"github.com/seed95/forward-proxy/internal/repo"
	"time"
)

type redisRepo struct {
	rdb        *redis.Client
	expiration time.Duration
}

var _ repo.CacheRepo = (*redisRepo)(nil)

func New(config *internal.RedisConfig) repo.CacheRepo {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Username: "", // TODO check load from config
		Password: config.Password,
		DB:       0, // TODO check load from config
	})
	return &redisRepo{
		rdb:        rdb,
		expiration: config.ExpirationTime,
	}
}

func (r *redisRepo) CacheResponse(ctx context.Context, url string, res interface{}) (err error) {
	return r.rdb.Set(ctx, url, res, r.expiration).Err()
}

func (r *redisRepo) GetCachedRequest(ctx context.Context, url string) (res interface{}, err error) {
	return r.rdb.Get(ctx, url).Result()
}

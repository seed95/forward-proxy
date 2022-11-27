package helper

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/seed95/forward-proxy/internal"
)

func NewRedisDatabase(config internal.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Username: "", // TODO check load from config
		Password: config.Password,
		DB:       0, // TODO check load from config
	})

	// Ping redis
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

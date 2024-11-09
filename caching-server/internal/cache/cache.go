package cache

import (
	"caching-server/internal/config"
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func GetRedis(cfg config.Config) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Address, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       0,
	})

	// Test connection by pinging Redis
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to redis server")
	}
	redisCache := RedisCache{}
	redisCache.Client = client

	return &redisCache, nil
}

func (rds *RedisCache) Clear() error {
	if err := rds.Client.FlushDB(context.TODO()).Err(); err != nil {
		return errors.Wrap(err, "unable to clear the cache")
	}
	return nil
}

package cache

import (
	"caching-server/internal/config"
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func GetRedis(cfg config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})

	// Test connection by pinging Redis
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to redis server")
	}

	return client, nil
}

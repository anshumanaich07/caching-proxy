package domain

import "context"

type CacheUsecase interface {
	Get(ctx context.Context, key string) (map[string]interface{}, error)
}

type CacheRepository interface {
	Get(ctx context.Context, key string) (map[string]interface{}, error)
}

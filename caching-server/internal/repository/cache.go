package repository

import "github.com/redis/go-redis/v9"

type cacheRepository struct {
	cache *redis.Client
}

func NewCacheRepository(c *redis.Client) cacheRepository {
	return cacheRepository{
		cache: c,
	}
}

func (repo cacheRepository) Get(key string) map[string]interface{} {
	return nil
}

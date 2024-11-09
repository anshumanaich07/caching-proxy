package repository

import (
	"caching-server/internal/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	cache  *redis.Client
	config *config.Config
}

func NewCacheRepository(c *redis.Client, sc *config.Config) cacheRepository {
	return cacheRepository{
		cache:  c,
		config: sc,
	}
}

func (repo cacheRepository) Get(ctx context.Context, key string) (map[string]interface{}, error) {
	cacheHit := true
	resMap := make(map[string]interface{})

	val, err := repo.cache.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			cacheHit = false
		} else {
			return nil, errors.Wrap(err, "unable to get value from redis")
		}
	}
	if val == "" {
		cacheHit = false
	} else {
		cacheHit = true
	}

	// cache miss
	if !cacheHit {
		api := fmt.Sprintf("%s:%d/%s/%s",
			repo.config.OriginHost,
			repo.config.OriginPort,
			"employee",
			key)
		fmt.Println("api called: ", api)
		res, err := http.Get(api)
		if err != nil {
			return nil, errors.Wrap(err, "unable to fetch data from main server")
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read response")
		}
		if err = json.Unmarshal([]byte(string(body)), &resMap); err != nil {
			return nil, errors.Wrap(err, "unable to unmarshal into map")
		}
		// set in cache
		bytes, err := json.Marshal(resMap)
		if err != nil {
			return nil, errors.Wrap(err, "unable to marshal")
		}
		if err = repo.cache.Set(ctx, key, bytes, 0).Err(); err != nil {
			return nil, errors.Wrap(err, "unable to set in cache")
		}
	} else { // cache hit
		err = json.Unmarshal([]byte(val), &resMap)
		if err != nil {
			return nil, errors.Wrap(err, "unable to unmarshal into map")
		}
	}

	return resMap, nil
}

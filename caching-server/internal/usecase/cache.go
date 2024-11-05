package usecase

import (
	"caching-server/internal/domain"
	"context"
)

type cacheUsecase struct {
	cacheRepo domain.CacheRepository
}

func NewCacheUsecase(repo domain.CacheRepository) cacheUsecase {
	return cacheUsecase{
		cacheRepo: repo,
	}
}

func (uc cacheUsecase) Get(ctx context.Context, key string) (map[string]interface{}, error) {
	// TODO
	// prepare the key for redis
	// key := fmt.Sprintf("")
	// check if exists
	// change the reponse, if needed
	return uc.cacheRepo.Get(ctx, key)
}

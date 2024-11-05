package usecase

import "caching-server/internal/domain"

type cacheUsecase struct {
	cacheRepo domain.CacheRepository
}

func NewCacheUsecase(repo domain.CacheRepository) cacheUsecase {
	return cacheUsecase{
		cacheRepo: repo,
	}
}

func (uc cacheUsecase) Get(key string) map[string]interface{} {
	return nil
}

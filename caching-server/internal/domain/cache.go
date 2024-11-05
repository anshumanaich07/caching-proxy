package domain

type CacheUsecase interface {
	Get(key string) map[string]interface{}
}

type CacheRepository interface {
	Get(key string) map[string]interface{}
}

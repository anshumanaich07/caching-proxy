package routes

import (
	"caching-server/internal/api/controllers"
	"caching-server/internal/cache"
	"caching-server/internal/config"
	"caching-server/internal/repository"
	"caching-server/internal/usecase"
	"net/http"
)

func InitRouter(client *cache.RedisCache, cfg *config.Config) *http.ServeMux {
	router := http.NewServeMux()

	repo := repository.NewCacheRepository(client, cfg)
	uc := usecase.NewCacheUsecase(repo)

	controller := controllers.NewCacheController(uc)

	router.HandleFunc("GET /employee/{id}", controller.GetEmployeeByID)

	return router
}

package routes

import (
	"caching-server/internal/api/controllers"
	"caching-server/internal/repository"
	"caching-server/internal/usecase"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func InitRouter(client *redis.Client) *http.ServeMux {
	router := http.NewServeMux()

	repo := repository.NewCacheRepository(client)
	uc := usecase.NewCacheUsecase(repo)

	controller := controllers.NewCacheController(uc)

	router.HandleFunc("GET /employee/{id}", controller.GetEmployeeByID)

	return router
}
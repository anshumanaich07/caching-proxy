package controllers

import (
	"caching-server/internal/domain"
	"net/http"
)

type cacheController struct {
	cacheUsecase domain.CacheUsecase
}

func NewCacheController(uc domain.CacheUsecase) cacheController {
	return cacheController{
		cacheUsecase: uc,
	}
}

func (controller cacheController) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
}

package controllers

import (
	"caching-server/internal/domain"
	"context"
	"fmt"
	"log"
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
	var ctx context.Context
	if r.Context() == nil {
		ctx = context.TODO()
	} else {
		ctx = r.Context()
	}

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing 'id parameter", http.StatusBadRequest)
		return
	}
	path := r.URL.Path

	ret, err := controller.cacheUsecase.Get(ctx, path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)

	// if not, forward the request to the main server
}

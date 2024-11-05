package main

import (
	"caching-server/internal/api/routes"
	"caching-server/internal/cache"
	"caching-server/internal/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the config
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	rds, err := cache.GetRedis(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	// router
	router := routes.InitRouter(rds)
	addr := fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port)
	log.Println("Caching server started...")
	log.Fatal(http.ListenAndServe(addr, router))
}

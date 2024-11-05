package main

import (
	"caching-server/internal/cache"
	"caching-server/internal/config"
	"fmt"
	"log"
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

	fmt.Println(rds)
}

package main

import (
	"caching-server/internal/api/routes"
	"caching-server/internal/cache"
	"caching-server/internal/config"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	originHost := flag.String("origin", "localhost", "origin address")
	port := flag.Int("port", 3000, "caching server port")
	flag.Parse()

	if originHost == nil || port == nil {
		log.Fatal("origin host or port cannot be empty")
	}

	// get the config
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg.OriginHost = *originHost
	cfg.OriginPort = *port

	rds, err := cache.GetRedis(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	// router
	router := routes.InitRouter(rds, cfg)
	addr := fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port)
	log.Println("Caching server started at - ", addr)
	log.Println("Origin server started at - ", fmt.Sprintf("%s:%d", cfg.OriginHost, cfg.OriginPort))
	log.Fatal(http.ListenAndServe(addr, router))
}

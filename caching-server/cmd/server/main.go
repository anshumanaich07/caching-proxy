package main

import (
	"caching-server/internal/api/routes"
	"caching-server/internal/cache"
	"caching-server/internal/config"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	originHost := flag.String("origin", "", "origin address")
	port := flag.Int("port", 0, "caching server port")
	clear := flag.Bool("clear", false, "clear the cache")
	flag.Parse()

	// get the config
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg.OriginHost = *originHost
	cfg.OriginPort = *port
	cfg.IsClear = *clear

	rds, err := cache.GetRedis(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	// check flags
	if cfg.IsClear && cfg.OriginHost == "" && cfg.OriginPort == 0 {
		if err = rds.Clear(); err != nil {
			log.Fatal(err)
		}
		log.Println("Cache cleared successfully")
		os.Exit(0)
	} else if cfg.IsClear && (cfg.OriginHost != "" || cfg.OriginPort != 0) {
		log.Fatal("cannot clear the cache and send request. Pick one")
	} else if !cfg.IsClear && (cfg.OriginHost == "" || cfg.OriginPort == 0) {
		log.Fatal("origin host or port cannot be empty")
	}

	// router
	router := routes.InitRouter(rds, cfg)
	addr := fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port)
	log.Println("Caching server started at - ", addr)
	log.Println("Origin server started at - ", fmt.Sprintf("%s:%d", cfg.OriginHost, cfg.OriginPort))
	log.Fatal(http.ListenAndServe(addr, router))
}

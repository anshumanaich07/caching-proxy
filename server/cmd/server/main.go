package main

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/api/routes"
	"server/internal/database"
	"server/pkg/config"
)

func main() {
	// get the env
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// database
	db, err := database.GetDatabase(*config)
	if err != nil {
		log.Fatal(err)
	}

	// init router
	router := routes.InitRouter(db)
	addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	log.Fatal(http.ListenAndServe(addr, router))
}

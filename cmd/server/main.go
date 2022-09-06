package main

import (
	config "bitcoind_rest_api/cmd/server/config"
	"bitcoind_rest_api/internal/routers"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	cfg, err := config.ParseConfig[config.Config]([]string{"./config/", "./cmd/server/config/"})
	if err != nil {
		log.Fatal("failed to parse config", "error", err)
	}

	routingService := routers.Initialize(cfg)
	routingService.RouteRequests()
}

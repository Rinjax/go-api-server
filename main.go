package main

import (
	"context"
	"jax/api"
	"jax/api/endpoints"
	"jax/config"
	"jax/hash"
	"jax/logger"
	"jax/store"
	"os"
	"os/signal"
	"syscall"
)



func main() {

	api, log := bootstrap()

	run(api, log)

	os.Exit(0)
}

func bootstrap() (*api.ApiServer, logger.Logger) {

	ctx := context.Background()

	// load the application config
	cfg, err := config.GetConfig()
    if err != nil {
        panic(err)
    }

	// build  dependancies
	log := logger.NewLogger(cfg)
	db := store.NewStore(cfg)
	hash := hash.NewHash(cfg)

	// build api server
	endp := endpoints.NewEndpoints(db, hash, log)
	apiSvr := api.NewServer(ctx, cfg, endp, log)

	return apiSvr, log
}

func run(api *api.ApiServer, log logger.Logger) {

	// Goroutine to start the HTTP server
	go api.Up()

	// Channel to listen for OS interrupt signals (Ctrl+C or SIGTERM)
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	// Wait for the termination signal (SIGINT or SIGTERM)
	<-stopCh
	log.Warn("received termination signal from OS")

	api.Down()
}
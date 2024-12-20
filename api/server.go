package api

import (
	"context"
	"jax/api/endpoints"
	"jax/config"
	"jax/logger"
	"net/http"
	"time"
)

type ApiServer struct {
	ctx context.Context
	config *config.ApiServerConfig
	log logger.Logger
	http *http.Server
}


func NewServer(ctx context.Context, config *config.Config, endpoints *endpoints.Endpoint, log logger.Logger) *ApiServer {
	s := &ApiServer{
		ctx: ctx,
		config: &config.Api,
		log: log,
	}

	s.http = &http.Server{
		Addr: s.config.Addr,
		Handler: routes(endpoints),
	}

	return s
}

func (s *ApiServer) Up() {
	s.log.Info("starting up api server...")
	s.log.Info("api server listening on:", s.config.Addr)

	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.log.Error("api server encountered a problem while serving:", err)
	}
}

func (s *ApiServer) Down() {
	s.log.Info("shutting down api server...")

	ctx, cancel := context.WithTimeout(s.ctx, time.Second*10)
    defer cancel()

    err := s.http.Shutdown(ctx)
	if err != nil {
		s.log.Error("api server failed to shutdown gracefully", err)
	} else {
		s.log.Info("api server shutdown gracefully")
	}
}
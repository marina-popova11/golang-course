package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marina-popova11/golang-course/task2/internal/gateway/handler"
	"github.com/marina-popova11/golang-course/task2/internal/gateway/usecase"

	"github.com/marina-popova11/golang-course/task2/internal/gateway/adapter/grpc_client"
	"github.com/marina-popova11/golang-course/task2/internal/gateway/config"
	swagger "github.com/swaggo/http-swagger"
)

func AppRun(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("config load: %w", err)
	}

	collectorClient, err := grpc_client.NewClient(ctx, cfg.Collector.Address)
	if err != nil {
		return fmt.Errorf("collector client: %w", err)
	}
	defer collectorClient.Close()

	uc := usecase.NewRepoUsecase(collectorClient)

	httpHandler := handler.NewHandler(uc)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	httpHandler.Register(router)

	if cfg.Swagger.Enabled {
		router.Get(cfg.Swagger.Path+"/*", swagger.WrapHandler)
	}

	server := &http.Server{
		Addr:         cfg.HTTP.Port,
		Handler:      router,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

	log.Printf("Gateway HTTP server starting on %s", cfg.HTTP.Port)
	log.Printf("Swagger available at http://localhost%s/index.html", cfg.Swagger.Path)

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("server listen: %w", err)
	}

	return nil
}

func main() {
	ctx := context.Background()
	if err := AppRun(ctx); err != nil {
		log.Fatalf("AppRun failed: %v", err)
	}
}

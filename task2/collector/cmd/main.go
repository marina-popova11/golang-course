package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/marina-popova11/golang-course/task2/collector/internal/adapter/github_client"
	"github.com/marina-popova11/golang-course/task2/collector/internal/config"
	"github.com/marina-popova11/golang-course/task2/collector/internal/handler"
	"github.com/marina-popova11/golang-course/task2/collector/internal/usecase"
	"google.golang.org/grpc"
)

func AppRun(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("Config load: %w", err)
	}

	githubClient := github_client.NewClient(cfg.GitHub.Token)
	uc := usecase.NewRepoUsecase(githubClient)
	grpcHandler := handler.NewHandler(uc)
	listener, err := net.Listen("tcp", cfg.GRPC.Port)
	if err != nil {
		return fmt.Errorf("Net listen: %w", err)
	}

	server := grpc.NewServer()
	grpcHandler.Register(server)
	log.Printf("Collector gRPC server starting on %s", cfg.GRPC.Port)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("Server serve: %w", err)
	}

	return nil
}

func main() {
	ctx := context.Background()
	if err := AppRun(ctx); err != nil {
		log.Fatalf("AppRun failed: %v", err)
	}
}

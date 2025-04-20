package main

import (
	"OrderProject/internal/service"
	test "OrderProject/pkg/api/test/api"
	"OrderProject/pkg/logger"
	"OrderProject/pkg/postgres"
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx := context.Background()
	ctx, _ = logger.New(ctx)

	postgresCfg := postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "root",
		Password: "qwerty",
		Database: "postgres",
	}

	db, err := postgres.New(postgresCfg)
	if err != nil {
		logger.GetLoggerFromContext(ctx).Fatal(ctx, "failed to connect to database", zap.Error(err))
		return
	}

	err = db.Ping(ctx)
	if err != nil {
		logger.GetLoggerFromContext(ctx).Fatal(ctx, "failed to ping database", zap.Error(err))
		return
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := service.New()
	server := grpc.NewServer()
	test.RegisterOrderServiceServer(server, srv)

	if err := server.Serve(lis); err != nil {
		logger.GetLoggerFromContext(ctx).Error(ctx, "failed to serve", zap.Error(err))
	}
}

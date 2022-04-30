package app

import (
	"context"
	"fmt"

	"go-microservice/internal/grpc"
	"go-microservice/internal/rest"
)

// StartServer : Starts both the gRPC and REST servers.
func StartServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go startGRPCServer(ctx, cancel)

	startRESTServer(ctx, cancel)

	<-ctx.Done()
	fmt.Println(" stopped all server..")
}

func startRESTServer(ctx context.Context, cancel context.CancelFunc) {
	s := rest.New()
	s.Start(ctx, cancel)
}

func startGRPCServer(ctx context.Context, cancel context.CancelFunc) {
	s := grpc.New()
	s.Start(ctx, cancel)
}

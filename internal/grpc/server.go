package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"go-microservice/internal/config"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	grpc *grpc.Server
}

func New() *Server {

	//can add interceptor middlewares in ChainUnaryServer
	return &Server{
		grpc: grpc.NewServer(
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
		),
	}
}

func (s *Server) Start(ctx context.Context, cancel context.CancelFunc) {

	err := registerService(s.grpc)
	if err != nil {
		fmt.Printf("error starting GRPC Server: %v \n", err.Error())
		cancel()
		return
	}

	go s.waitForShutdownGRPC(cancel)

	port := strconv.Itoa(config.GRPCServer.Port)

	fmt.Printf("starting GRPC server on port %s... \n", strconv.Itoa(config.GRPCServer.Port))
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("error starting GRPC Server: %v \n", err.Error())
		cancel()
		return
	}

	if err = s.grpc.Serve(lis); err != nil {
		cancel()
	}
}

func (s *Server) waitForShutdownGRPC(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-c
	fmt.Println("shutting down gRPC server...")
	s.grpc.GracefulStop()
	cancel() // call the cancelFunc to close the shared interrupt channel between REST and gRPC and shutdown both servers
}

func registerService(s *grpc.Server) error {
	registerHealthCheck(s)
	return nil
}

func registerHealthCheck(s *grpc.Server) {
	server := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, server)
	server.SetServingStatus("Zalango Service", grpc_health_v1.HealthCheckResponse_SERVING)
}

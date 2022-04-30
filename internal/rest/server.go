package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"go-microservice/internal/config"
	"go-microservice/internal/flagr"
	flagrHandler "go-microservice/internal/handler/v1/flagr"

	"github.com/checkr/goflagr"
)

type Server struct {
	server *http.Server
}

func New() *Server {

	client := goflagr.NewAPIClient(&config.Flagr)
	flagrService := flagr.NewService(client)
	flagrHandler := flagrHandler.New(flagrService)

	handler := router(flagrHandler)
	server := &Server{
		server: &http.Server{
			Addr:         ":" + strconv.Itoa(config.RESTServer.Port),
			Handler:      handler,
			TLSConfig:    nil,
			ReadTimeout:  config.RESTServer.ReadTimeout,
			WriteTimeout: config.RESTServer.WriteTimeout,
		}}

	return server
}

func (s *Server) Start(ctx context.Context, cancel context.CancelFunc) {
	go s.waitForShutDown(ctx, cancel)

	fmt.Printf("starting REST server on port %s... \n", strconv.Itoa(config.RESTServer.Port))
	go func() {
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("failed to start the server : " + err.Error())
			cancel()
		}
	}()
}

func (s *Server) waitForShutDown(ctx context.Context, cancel context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	fmt.Println("stopping REST server")
	err := s.server.Shutdown(ctx)
	if err != nil {
		fmt.Printf("error in closing down server gracefully %s", err.Error())
	}
	cancel()
}

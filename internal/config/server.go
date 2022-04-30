package config

import (
	"time"
)

type RESTServerConfig struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Port         int
}

type GRPCServerConfig struct {
	RequestTimeout time.Duration
	Port           int
}

var RESTServer RESTServerConfig
var GRPCServer GRPCServerConfig

func initRESTServerConfig() {
	RESTServer = RESTServerConfig{
		ReadTimeout:  mustGetDurationMs("REST_SERVER_READ_TIMEOUT_MS"),
		WriteTimeout: mustGetDurationMs("REST_SERVER_WRITE_TIMEOUT_MS"),
		Port:         mustGetInt("REST_SERVER_PORT"),
	}
}

func initGRPCServerConfig() {
	GRPCServer = GRPCServerConfig{
		RequestTimeout: mustGetDurationMs("GRPC_REQUEST_TIMEOUT_MS"),
		Port:           mustGetInt("GRPC_SERVER_PORT"),
	}
}

package server

import (
	pkgGrpc "github.com/vogiaan1904/order-svc/pkg/grpc"
	pkgLog "github.com/vogiaan1904/order-svc/pkg/log"
	"github.com/vogiaan1904/order-svc/pkg/mongo"
	"go.temporal.io/sdk/client"
)

type server struct {
	l        pkgLog.Logger
	port     string
	db       mongo.Database
	grpc     pkgGrpc.Clients
	temporal client.Client
}

type Config struct {
	Port     string
	Db       mongo.Database
	Grpc     pkgGrpc.Clients
	Temporal client.Client
}

func New(l pkgLog.Logger, cfg Config) *server {
	return &server{
		l:        l,
		port:     cfg.Port,
		db:       cfg.Db,
		grpc:     cfg.Grpc,
		temporal: cfg.Temporal,
	}
}

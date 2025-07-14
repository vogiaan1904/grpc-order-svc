package main

import (
	"github.com/vogiaan1904/order-svc/config"
	"github.com/vogiaan1904/order-svc/internal/appconfig/mongo"
	"github.com/vogiaan1904/order-svc/internal/appconfig/temporal"
	"github.com/vogiaan1904/order-svc/internal/server"
	pkgGrpc "github.com/vogiaan1904/order-svc/pkg/grpc"
	pkgLog "github.com/vogiaan1904/order-svc/pkg/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Log.Level,
		Encoding: cfg.Log.Encoding,
		Mode:     cfg.Log.Mode,
	})

	// MongoDB connection
	mongoCli, err := mongo.Connect(cfg.Mongo.DatabaseUri)
	if err != nil {
		panic(err)
	}
	defer mongo.Disconnect(mongoCli)
	db := mongoCli.Database(cfg.Mongo.DatabaseName)

	// Temporal connection
	temporalCli, err := temporal.Connect(cfg.Temporal)
	if err != nil {
		panic(err)
	}
	defer temporal.Disconnect(temporalCli)

	// gRPC clients connection
	grpcClis, cleanupGrpc, err := pkgGrpc.InitClients(cfg.Grpc.ProductSvcAddr, l, cfg.Log.RedactFields)
	if err != nil {
		panic(err)
	}
	defer cleanupGrpc()

	s := server.New(l, server.Config{
		Port:     cfg.Server.Port,
		Db:       db,
		Grpc:     grpcClis,
		Temporal: temporalCli,
	})
	if err := s.Run(); err != nil {
		panic(err)
	}
}

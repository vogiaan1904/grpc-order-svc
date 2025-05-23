package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/vogiaan1904/order-svc/config"
	"github.com/vogiaan1904/order-svc/internal/appconfig/mongo"
	"github.com/vogiaan1904/order-svc/internal/interceptors"
	repository "github.com/vogiaan1904/order-svc/internal/repositories"
	service "github.com/vogiaan1904/order-svc/internal/services"
	grpcservices "github.com/vogiaan1904/order-svc/pkg/grpc"
	pkgLog "github.com/vogiaan1904/order-svc/pkg/log"
	order "github.com/vogiaan1904/order-svc/protogen/golang/order"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Log.Level,
		Encoding: cfg.Log.Encoding,
		Mode:     cfg.Log.Mode,
	})

	// Initialize Temporal Client
	tCli, err := client.Dial(client.Options{
		HostPort:  cfg.Temporal.HostPort,
		Namespace: cfg.Temporal.Namespace,
	})
	if err != nil {
		l.Fatal(context.Background(), "Unable to create Temporal Client", "error", err)
	}
	defer tCli.Close()
	l.Info(context.Background(), "Temporal Client connected.")

	const addr = "127.0.0.1:50054"
	lnr, err := net.Listen("tcp", addr)
	if err != nil {
		l.Fatal(context.Background(), "failed to listen", "error", err)
	}

	// MongoDB connection
	mCli, err := mongo.Connect(cfg.Mongo.DatabaseUri)
	if err != nil {
		l.Fatal(context.Background(), "Failed to connect to MongoDB", "error", err)
	}
	defer mongo.Disconnect(mCli)
	db := mCli.Database(cfg.Mongo.DatabaseName)

	// gRPC clients
	grpcClis, cleanupGrpc, err := grpcservices.InitGrpcClients(cfg.Grpc.ProductSvcAddr, l, cfg.Log.RedactFields)
	if err != nil {
		l.Fatal(context.Background(), "failed to initialize gRPC clients", "error", err)
	}
	defer cleanupGrpc()

	orderRepo := repository.NewOrderRepository(l, db)
	orderSvc := service.NewOrderService(l, orderRepo, grpcClis.Product, tCli)

	// gRPC server
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors.ValidationInterceptor, interceptors.ErrorHandlerInterceptor),
	)

	order.RegisterOrderServiceServer(server, orderSvc)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		l.Info(context.Background(), "gRPC server started", "address", addr)
		if err := server.Serve(lnr); err != nil {
			l.Fatal(context.Background(), "failed to serve gRPC", "error", err)
		}
	}()

	<-sigCh
	l.Info(context.Background(), "Shutting down gRPC server...")

	server.GracefulStop()
	l.Info(context.Background(), "Server stopped")
}

package main

import (
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
	"github.com/vogiaan1904/order-svc/pkg/grpcservices"
	pkgLog "github.com/vogiaan1904/order-svc/pkg/log"
	order "github.com/vogiaan1904/order-svc/protogen/golang/order"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Initialize logger
	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Log.Level,
		Encoding: cfg.Log.Encoding,
		Mode:     cfg.Log.Mode,
	})

	const addr = "127.0.0.1:50054"
	lnr, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// MongoDB connection
	mClient, err := mongo.Connect(cfg.Mongo.DatabaseUri)
	if err != nil {
		panic(err)
	}
	defer mongo.Disconnect(mClient)
	db := mClient.Database(cfg.Mongo.DatabaseName)

	// gRPC clients
	grpcClients, cleanupGrpc, err := grpcservices.InitGrpcClients(cfg.Grpc.ProductSvcAddr, l, cfg.Log.RedactFields)
	if err != nil {
		log.Fatalf("failed to initialize gRPC clients: %v", err)
	}
	defer cleanupGrpc()

	orderRepo := repository.NewOrderRepository(l, db)
	orderSvc := service.NewOrderService(l, orderRepo, grpcClients.Product)

	// gRPC server
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors.ValidationInterceptor, interceptors.ErrorHandlerInterceptor),
	)

	order.RegisterOrderServiceServer(server, orderSvc)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("gRPC server started on %s", addr)
		if err := server.Serve(lnr); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-sigCh
	log.Println("Shutting down gRPC server...")

	server.GracefulStop()
	log.Println("Server stopped")
}

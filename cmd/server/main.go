package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/vogiaan1904/order-svc/config"
	"github.com/vogiaan1904/order-svc/internal/appconfig/mongo"
	repository "github.com/vogiaan1904/order-svc/internal/repositories"
	"github.com/vogiaan1904/order-svc/internal/service"
	order "github.com/vogiaan1904/order-svc/protogen/golang/order"
	product "github.com/vogiaan1904/order-svc/protogen/golang/product"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	const addr = "127.0.0.1:50054"
	lnr, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	productSvcAddr := cfg.Grpc.ProductServiceAddress
	productCnn, err := grpc.Dial(productSvcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to product service: %v", err)
	}
	defer productCnn.Close()
	productClient := product.NewProductServiceClient(productCnn)

	mClient, err := mongo.Connect(cfg.Mongo.DatabaseUri)
	if err != nil {
		panic(err)
	}
	defer mongo.Disconnect(mClient)

	db := mClient.Database(cfg.Mongo.DatabaseName)
	orderRepo := repository.NewOrderRepository(db)
	orderSvc := service.NewOrderService(orderRepo, productClient)

	order.RegisterOrderServiceServer(server, orderSvc)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("gRPC server started  on %s", addr)
		if err := server.Serve(lnr); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-sigCh
	log.Println("Shutting down gRPC server...")

	server.GracefulStop()
	log.Println("Server stopped")
}

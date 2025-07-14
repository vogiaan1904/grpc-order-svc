package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	orderGRPC "github.com/vogiaan1904/order-svc/internal/delivery/grpc"
	"github.com/vogiaan1904/order-svc/internal/interceptors"
	orderMongo "github.com/vogiaan1904/order-svc/internal/repositories/mongo"
	"github.com/vogiaan1904/order-svc/internal/services"
	orderpb "github.com/vogiaan1904/order-svc/protogen/golang/order"
	"google.golang.org/grpc"
)

func (s *server) Run() error {
	orderRepo := orderMongo.NewOrderRepo(s.l, s.db)
	orderSvc := services.NewOrderService(s.l, orderRepo, s.grpc.Product, s.temporal)
	orderSrv := orderGRPC.NewOrderServiceServer(s.l, orderSvc)

	lnr, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors.ValidationInterceptor, interceptors.ErrorHandlerInterceptor),
	)
	orderpb.RegisterOrderServiceServer(grpcServer, orderSrv)

	go func() {
		log.Printf("gRPC server is listening on port: %s", s.port)
		if err := grpcServer.Serve(lnr); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	sig := <-sigCh
	log.Printf("Received signal %v, shutting down gracefully...", sig)

	grpcServer.GracefulStop()
	log.Printf("Server stopped")

	return nil
}

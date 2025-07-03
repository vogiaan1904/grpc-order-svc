package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/vogiaan1904/order-svc/config"
	order "github.com/vogiaan1904/order-svc/internal/handlers/grpc"
	"github.com/vogiaan1904/order-svc/internal/interceptors"
	orderMongoRepo "github.com/vogiaan1904/order-svc/internal/repositories/mongo"
	"github.com/vogiaan1904/order-svc/internal/services"
	pkgGrpc "github.com/vogiaan1904/order-svc/pkg/grpc"
	pkgLog "github.com/vogiaan1904/order-svc/pkg/log"
	"github.com/vogiaan1904/order-svc/pkg/mongo"
	orderGRPCService "github.com/vogiaan1904/order-svc/protogen/golang/order"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc"
)

type server struct {
	l        pkgLog.Logger
	cfg      *config.Config
	db       mongo.Database
	grpcClis *pkgGrpc.Clients
	tprCli   client.Client
}

func NewServer(l pkgLog.Logger, cfg *config.Config, db mongo.Database, grpcClis *pkgGrpc.Clients, tprCli client.Client) *server {
	return &server{l: l, cfg: cfg, db: db, grpcClis: grpcClis, tprCli: tprCli}
}

func (s *server) Run() error {
	// ctx := context.Background()

	orderRepo := orderMongoRepo.NewOrderRepository(s.l, s.db)
	orderSvc := services.NewOrderService(s.l, orderRepo, s.grpcClis.Product, s.tprCli)

	addr := fmt.Sprintf(":%s", s.cfg.Server.Port)
	lnr, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors.ValidationInterceptor, interceptors.ErrorHandlerInterceptor),
	)

	orderService := order.NewOrderService(s.l, orderSvc)
	orderGRPCService.RegisterOrderServiceServer(grpcServer, orderService)

	go func() {
		log.Printf("gRPC server is listening on port: %s", s.cfg.Server.Port)
		if err := grpcServer.Serve(lnr); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// TODO: Add http server
	// go func() {
	//     s.l.Info(ctx, "HTTP server starting...")
	//     if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//         s.l.Fatal(ctx, "HTTP server failed", "error", err)
	//     }
	// }()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	sig := <-sigCh
	log.Printf("Received signal %v, shutting down gracefully...", sig)

	grpcServer.GracefulStop()
	log.Println("Server stopped")

	return nil
}

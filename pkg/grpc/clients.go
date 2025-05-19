package grpcservices

import (
	"context"
	"log"

	"github.com/vogiaan1904/order-svc/internal/interceptors"
	pkgLog "github.com/vogiaan1904/order-svc/pkg/log"
	"github.com/vogiaan1904/order-svc/protogen/golang/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Product product.ProductServiceClient
}

type cleanupFunc func()

func InitGrpcClients(productAddr string, logger pkgLog.Logger, redactedFields []string) (*GrpcClients, cleanupFunc, error) {
	// Create a background context for logging
	ctx := context.Background()

	var cleanupFuncs []cleanupFunc
	clients := &GrpcClients{}

	// Create interceptor with logger and redacted fields
	loggingInterceptor := interceptors.GrpcClientLoggingInterceptor(logger, redactedFields)

	// Create client connection with the interceptor
	productConn, err := grpc.NewClient(
		productAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(loggingInterceptor),
	)
	if err != nil {
		return nil, nil, err
	}
	clients.Product = product.NewProductServiceClient(productConn)

	// Add more clients here...

	cleanupFuncs = append(cleanupFuncs, func() {
		if err := productConn.Close(); err != nil {
			log.Printf("failed to close product gRPC connection: %v", err)
		}
	})

	cleanupFunc := func() {
		for _, fn := range cleanupFuncs {
			fn()
		}
		logger.Info(ctx, "gRPC clients cleaned up")
	}

	logger.Info(ctx, "gRPC clients initialized")

	return clients, cleanupFunc, nil
}

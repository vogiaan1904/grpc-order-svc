package grpc

import (
	"log"

	"github.com/vogiaan1904/order-svc/internal/interceptors"
	pkgLog "github.com/vogiaan1904/order-svc/pkg/log"
	"github.com/vogiaan1904/order-svc/protogen/golang/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Product product.ProductServiceClient
}

type cleanupFunc func()

func InitClients(productAddr string, logger pkgLog.Logger, redactedFields []string) (*Clients, cleanupFunc, error) {
	var cleanupFuncs []cleanupFunc
	clients := &Clients{}

	loggingInterceptor := interceptors.GrpcClientLoggingInterceptor(logger, redactedFields)

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
			log.Printf("Failed to close product gRPC connection: %v", err)
		}
	})

	cleanupFunc := func() {
		for _, fn := range cleanupFuncs {
			fn()
		}
		log.Println("gRPC clients cleaned up.")
	}

	log.Println("gRPC clients initialized!")

	return clients, cleanupFunc, nil
}

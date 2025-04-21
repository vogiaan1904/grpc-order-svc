package grpcservices

import (
	"log"

	"github.com/vogiaan1904/order-svc/protogen/golang/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Product product.ProductServiceClient
}

type cleanupFunc func()

func InitGrpcClients(productAddr string) (*GrpcClients, cleanupFunc, error) {
	var cleanupFuncs []cleanupFunc
	clients := &GrpcClients{}

	productConn, err := grpc.NewClient(productAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		log.Println("gRPC clients cleaned up")
	}

	log.Println("gRPC clients initialized")

	return clients, cleanupFunc, nil
}

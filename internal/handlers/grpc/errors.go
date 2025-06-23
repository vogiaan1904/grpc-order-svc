package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	services "github.com/vogiaan1904/order-svc/internal/services"
)

var (
	ErrOrderNotFound   = status.Error(codes.NotFound, "Order not found")
	ErrProductNotFound = status.Error(codes.NotFound, "Product not found")
)

var (
	ErrInvalidOrderData  = status.Error(codes.InvalidArgument, "Invalid order data")
	ErrInsufficientStock = status.Error(codes.InvalidArgument, "Insufficient stock")
	ErrOrderAlreadyPaid  = status.Error(codes.InvalidArgument, "Order already paid")
)

func mapGRPCErrorCode(err error) error {
	if err == nil {
		return nil
	}

	switch err {
	case services.ErrOrderNotFound:
		return ErrOrderNotFound
	case services.ErrInvalidOrderData:
		return ErrInvalidOrderData
	case services.ErrProductNotFound:
		return ErrProductNotFound
	case services.ErrInsufficientStock:
		return ErrInsufficientStock
	case services.ErrOrderAlreadyPaid:
		return ErrOrderAlreadyPaid
	}

	return status.Error(codes.Internal, "Internal server error")
}

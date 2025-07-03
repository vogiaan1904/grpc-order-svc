package services

import (
	"context"

	"github.com/vogiaan1904/order-svc/internal/models"
)

type OrderService interface {
	CreateOrder(ctx context.Context, input CreateOrderInput) (CreateOrderOutput, error)
	FindOneOrder(ctx context.Context, input GetOneOrderInput) (models.Order, error)
	GetOrders(ctx context.Context, input GetOrdersInput) (GetOrdersOutput, error)
	UpdateOrderStatus(ctx context.Context, input UpdateOrderStatusInput) error
}

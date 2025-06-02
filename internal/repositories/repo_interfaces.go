package repository

import (
	"context"

	"github.com/vogiaan1904/order-svc/internal/models"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, opt CreateOrderOptions) (models.Order, error)
	FindManyOrders(ctx context.Context, opt FindManyOrderOptions) ([]models.Order, error)
	FindOneOrder(ctx context.Context, opt FindOneOrderOptions) (models.Order, error)
	UpdateOrder(ctx context.Context, o models.Order) error
}

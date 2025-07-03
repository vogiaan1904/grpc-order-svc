package repositories

import (
	"context"

	"github.com/vogiaan1904/order-svc/internal/models"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, opts CreateOrderOptions) (models.Order, error)
	GetOrders(ctx context.Context, opts GetOrdersOptions) ([]models.Order, error)
	FindOneOrder(ctx context.Context, opts FindOneOrderOptions) (models.Order, error)
	UpdateOrder(ctx context.Context, o models.Order) error
}

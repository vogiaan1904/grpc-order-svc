package repository

import (
	"github.com/vogiaan1904/order-svc/internal/models"
)

type CreateOrderOptions struct {
	Code        string
	UserID      string
	Items       []models.OrderItem
	TotalAmount float64
	Status      models.OrderStatus
}

type GetOrdersFilter struct {
	Code   string
	UserID string
	Status string
}

type GetOrdersOptions struct {
	GetOrdersFilter
}

type FindOneOrderOptions struct {
	GetOrdersFilter
	ID string
}

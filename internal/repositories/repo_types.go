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

type FindFilter struct {
	Code   string
	UserID string
	Status *models.OrderStatus
}

type FindManyOrderOptions struct {
	FindFilter
}

type FindOneOrderOptions struct {
	FindFilter
	ID string
}

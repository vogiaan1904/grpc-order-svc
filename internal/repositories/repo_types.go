package repository

import "github.com/vogiaan1904/order-svc/internal/models"

type CreateOrderOptions struct {
	UserID       string
	ProductID    string
	ProductName  string
	ProductPrice float64
	Quantity     int32
	TotalAmount  float64
	Status       models.OrderStatus
}

type FindFilter struct {
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

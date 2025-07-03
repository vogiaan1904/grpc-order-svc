package services

import (
	"github.com/vogiaan1904/order-svc/internal/models"
)

type OrderWorkflowParams struct {
	OrderCode   string
	UserId      string
	TotalAmount float64
}

type CreateOrderInput struct {
	UserID      string
	Items       []models.OrderItem
	TotalAmount float64
}

type CreateOrderOutput struct {
	OrderCode  string
	WorkflowID string
	PaymentURL string
}

type GetOneOrderInput struct {
	ID   string
	Code string
}

type GetOrdersInput struct {
	UserID string
	Status string
}

type GetOrdersOutput struct {
	Orders []models.Order
	// Paginator paginator.Paginator
}

type UpdateOrderStatusInput struct {
	ID     string
	Code   string
	Status string
}

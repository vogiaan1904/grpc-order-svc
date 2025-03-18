package models

import "time"

type OrderStatus string

const (
    OrderStatusPending   OrderStatus = "PENDING"
    OrderStatusCompleted OrderStatus = "COMPLETED"
    OrderStatusCancelled OrderStatus = "CANCELLED"
)

type Order struct {
    ID        string      `bson:"_id,omitempty"`
    UserID    string      `bson:"user_id"`
    ProductID string      `bson:"product_id"`
    Quantity  int32       `bson:"quantity"`
    TotalPrice float64    `bson:"total_price"`
    Status    OrderStatus `bson:"status"`
    CreatedAt time.Time   `bson:"created_at"`
    UpdatedAt time.Time   `bson:"updated_at"`
}
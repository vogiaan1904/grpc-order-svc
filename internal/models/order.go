package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "PENDING"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusCompleted  OrderStatus = "COMPLETED"
	OrderStatusCancelled  OrderStatus = "CANCELLED"
)

type Order struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       string             `bson:"user_id"`
	ProductID    string             `bson:"product_id"`
	ProductName  string             `bson:"product_name"`
	ProductPrice float64            `bson:"product_price"`
	Quantity     int32              `bson:"quantity"`
	TotalAmount  float64            `bson:"total_price"`
	Status       OrderStatus        `bson:"status"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	DeletedAt    *time.Time         `bson:"deleted_at,omitempty"`
}

package service

import "errors"

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrOrderNotFound     = errors.New("order not found")
	ErrInvalidOrderData  = errors.New("invalid order data")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrOrderAlreadyPaid  = errors.New("order already paid")
)

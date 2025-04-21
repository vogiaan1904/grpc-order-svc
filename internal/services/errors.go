package service

import "errors"

var WarnErrors = []error{
	ErrInvalidInput,
	ErrRequiredField,
	ErrProductNotFound,
	ErrProductOutOfStock,
}

var (
	ErrInvalidInput      = errors.New("invalid input")
	ErrRequiredField     = errors.New("required field is missing")
	ErrProductNotFound   = errors.New("product not found")
	ErrProductOutOfStock = errors.New("product out of stock")
)

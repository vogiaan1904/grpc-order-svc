package order

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var WarnErrors = []error{
	ErrInvalidInput,
	ErrRequiredField,
}

var (
	ErrRequiredField = errors.New("required field is missing")
	ErrInvalidInput  = errors.New("invalid input")
)

func (r *CreateRequest) Validate() error {
	if r.UserId == "" {
		return ErrRequiredField
	}
	if r.ProductId == "" {
		return ErrRequiredField
	} 
	if r.Quantity <= 0 {
		return ErrInvalidInput
	}

	return nil
}

func (r *FindOneRequest) Validate() error {
	if r.Id == "" {
		return ErrRequiredField
	}
	if _, err := primitive.ObjectIDFromHex(r.Id); err != nil {
		return ErrInvalidInput
	}

	return nil
}

func IsValidOrderStatus(status OrderStatus) bool {
	_, ok := OrderStatus_name[int32(status)]
	return ok
}

func (r *FindManyRequest) Validate() error {
	if !IsValidOrderStatus(r.Status) && r.Status != 0 {
		return ErrInvalidInput
	}
	return nil
}

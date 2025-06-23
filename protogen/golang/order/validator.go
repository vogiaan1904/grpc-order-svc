package order

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CreateRequest) Validate() bool {
	if r.UserId == "" {
		return false
	}

	if len(r.Items) == 0 {
		return false
	}

	for _, item := range r.Items {
		if item.ProductId == "" {
			return false
		}
		if item.Quantity <= 0 {
			return false
		}
	}

	return true
}

func (r *FindOneRequest) Validate() bool {
	if r.GetId() == "" && r.GetCode() == "" {
		return false
	}

	if r.GetId() != "" {
		if _, err := primitive.ObjectIDFromHex(r.GetId()); err != nil {
			return false
		}
	}

	return true
}

func IsValidOrderStatus(status OrderStatus) bool {
	_, ok := OrderStatus_name[int32(status)]

	return ok
}

func (r *FindManyRequest) Validate() bool {
	if r.Status != OrderStatus_ORDER_STATUS_UNSPECIFIED {
		if !IsValidOrderStatus(r.Status) {
			return false
		}
	}

	if r.UserId == "" {
		return false
	}

	return true
}

func (r *UpdateStatusRequest) Validate() bool {
	if r.GetId() == "" && r.GetCode() == "" {
		return false
	}

	if r.Status == OrderStatus_ORDER_STATUS_UNSPECIFIED {
		return false
	}

	return true
}

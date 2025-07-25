package mongo

import (
	"github.com/vogiaan1904/order-svc/internal/models"
	repo "github.com/vogiaan1904/order-svc/internal/repositories"
)

func (r *implOrderRepository) buildOrderModel(opt repo.CreateOrderOptions) models.Order {
	now := r.clock()
	return models.Order{
		ID:          r.db.NewObjectID(),
		Code:        opt.Code,
		UserID:      opt.UserID,
		Items:       opt.Items,
		TotalAmount: opt.TotalAmount,
		Status:      opt.Status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

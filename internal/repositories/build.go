package repository

import "github.com/vogiaan1904/order-svc/internal/models"

func (r *implOrderRepository) buildOrderModel(opt CreateOrderOptions) models.Order {
	now := r.clock()

	return models.Order{
		ID:          r.db.NewObjectID(),
		UserID:      opt.UserID,
		Items:       opt.Items,
		TotalAmount: opt.TotalAmount,
		Status:      opt.Status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

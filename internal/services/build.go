package service

import (
	"github.com/vogiaan1904/order-svc/internal/models"
	repository "github.com/vogiaan1904/order-svc/internal/repositories"
	order "github.com/vogiaan1904/order-svc/protogen/golang/order"
)

func protoOrderStatusToPtr(status order.OrderStatus) *models.OrderStatus {
	if status == order.OrderStatus_DEFAULT {
		return nil
	}
	str, ok := order.OrderStatus_name[int32(status)]
	if !ok {
		return nil
	}
	s := models.OrderStatus(str)
	return &s
}

func (svc *implOrderService) buildFindOptions(req *order.FindManyRequest) repository.FindManyOrderOptions {
	return repository.FindManyOrderOptions{
		FindFilter: repository.FindFilter{
			UserID: req.UserId,
			Status: protoOrderStatusToPtr(req.Status),
		},
	}
}

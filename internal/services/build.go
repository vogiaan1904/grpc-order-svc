package service

import (
	"github.com/vogiaan1904/order-svc/internal/models"
	repository "github.com/vogiaan1904/order-svc/internal/repositories"
	order "github.com/vogiaan1904/order-svc/protogen/golang/order"
)

func (svc *implOrderService) buildFindOptions(req *order.FindManyRequest) repository.FindManyOrderOptions {
	return repository.FindManyOrderOptions{
		FindFilter: repository.FindFilter{
			UserID: req.UserId,
			Status: svc.protoOrderStatusToPtr(req.Status),
		},
	}
}

func (svc *implOrderService) buildCreateOptions(req *order.CreateRequest, items []models.OrderItem) repository.CreateOrderOptions {

	// handle build///

	return repository.CreateOrderOptions{
		UserID: req.UserId,
		Items:  items,
	}
}

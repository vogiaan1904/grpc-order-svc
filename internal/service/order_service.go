package service

import (
	"context"

	"github.com/vogiaan1904/order-svc/internal/models"
	repository "github.com/vogiaan1904/order-svc/internal/repositories"
	order "github.com/vogiaan1904/order-svc/protogen/golang/order"
	"github.com/vogiaan1904/order-svc/protogen/golang/product"
	"google.golang.org/protobuf/types/known/emptypb"
)

type implOrderService struct {
	repo       repository.OrderRepository
	productSvc product.ProductServiceClient
	order.UnimplementedOrderServiceServer
}

func NewOrderService(repo repository.OrderRepository, productSvc product.ProductServiceClient) order.OrderServiceServer {
	return &implOrderService{
		repo:       repo,
		productSvc: productSvc,
	}
}

func (svc *implOrderService) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*emptypb.Empty, error) {
	var p *product.ProductData
	res, err := svc.productSvc.FindById(ctx, &product.FindByIdRequest{
		Id: req.ProductId,
	})
	if err != nil {
		return nil, err
	}

	if res != nil {
		p = res.Product
	}

	if p == nil {
		return nil, err
	}

	if p.Stock < req.Quantity {
		return nil, err
	}
	_, err = svc.repo.CreateOrder(ctx, repository.CreateOrderOptions{
		UserID:       req.UserId,
		ProductID:    req.ProductId,
		ProductName:  p.Name,
		ProductPrice: p.Price,
		Quantity:     req.Quantity,
		TotalAmount:  p.Price * float64(req.Quantity),
		Status:       models.OrderStatusPending,
	})
	if err != nil {
		return nil, err
	}

	// Update product stock

	return &emptypb.Empty{}, nil
}

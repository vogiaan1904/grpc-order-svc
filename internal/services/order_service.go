package service

import (
	"context"

	"github.com/vogiaan1904/order-svc/internal/models"
	repository "github.com/vogiaan1904/order-svc/internal/repositories"
	"github.com/vogiaan1904/order-svc/pkg/log"
	"github.com/vogiaan1904/order-svc/pkg/mongo"
	order "github.com/vogiaan1904/order-svc/protogen/golang/order"
	"github.com/vogiaan1904/order-svc/protogen/golang/product"
	"google.golang.org/protobuf/types/known/emptypb"
)

type implOrderService struct {
	l          log.Logger
	repo       repository.OrderRepository
	productSvc product.ProductServiceClient
	order.UnimplementedOrderServiceServer
}

func NewOrderService(l log.Logger, repo repository.OrderRepository, productSvc product.ProductServiceClient) order.OrderServiceServer {
	return &implOrderService{
		l:          l,
		repo:       repo,
		productSvc: productSvc,
	}
}

func (svc *implOrderService) Create(ctx context.Context, req *order.CreateRequest) (*emptypb.Empty, error) {
	var p *product.ProductData

	// This gRPC client call will be logged by the GrpcClientLoggingInterceptor
	res, err := svc.productSvc.FindById(ctx, &product.FindByIdRequest{
		Id: req.ProductId,
	})
	if err != nil {
		svc.l.Errorf(ctx, "Error FindById: %v", err)
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

	_, err = svc.productSvc.DecreaseStock(ctx, &product.DecreaseStockRequest{
		Id:       req.ProductId,
		Quantity: req.Quantity,
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (svc *implOrderService) FindOne(ctx context.Context, req *order.FindOneRequest) (*order.FindOneResponse, error) {
	o, err := svc.repo.FindOneOrder(ctx, repository.FindOneOrderOptions{
		ID:         req.Id,
		FindFilter: repository.FindFilter{},
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			svc.l.Warnf(ctx, "Order not found: %v", err)
			return nil, nil
		}
		svc.l.Errorf(ctx, "Error FindOneOrder: %v", err)
		return nil, err
	}

	return &order.FindOneResponse{
		Order: &order.OrderData{
			Id:           o.ID.Hex(),
			UserId:       o.UserID,
			ProductId:    o.ProductID,
			ProductName:  o.ProductName,
			ProductPrice: o.ProductPrice,
			Quantity:     o.Quantity,
			TotalAmount:  o.TotalAmount,
			Status:       order.OrderStatus(order.OrderStatus_value[string(o.Status)]),
		},
	}, nil
}

func (svc *implOrderService) FindMany(ctx context.Context, req *order.FindManyRequest) (*order.FindManyResponse, error) {
	opt := svc.buildFindOptions(req)
	os, err := svc.repo.FindManyOrders(ctx, opt)
	if err != nil {
		svc.l.Errorf(ctx, "Error FindManyOrders: %v", err)
		return nil, err
	}
	svc.l.Debugf(ctx, "orders: %v", os)

	var orders []*order.OrderData
	for _, o := range os {
		orders = append(orders, &order.OrderData{
			Id:           o.ID.Hex(),
			UserId:       o.UserID,
			ProductId:    o.ProductID,
			ProductName:  o.ProductName,
			ProductPrice: o.ProductPrice,
			Quantity:     o.Quantity,
			TotalAmount:  o.TotalAmount,
			Status:       order.OrderStatus(order.OrderStatus_value[string(o.Status)]),
		})
	}

	return &order.FindManyResponse{
		Orders: orders,
	}, nil
}

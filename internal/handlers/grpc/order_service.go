package grpc

import (
	"context"

	"github.com/vogiaan1904/order-svc/internal/models"
	services "github.com/vogiaan1904/order-svc/internal/services"
	"github.com/vogiaan1904/order-svc/pkg/log"
	orderpb "github.com/vogiaan1904/order-svc/protogen/golang/order"
	"google.golang.org/protobuf/types/known/emptypb"
)

type orderService struct {
	l   log.Logger
	svc services.OrderService
	orderpb.UnimplementedOrderServiceServer
}

func NewOrderService(l log.Logger, svc services.OrderService) orderpb.OrderServiceServer {
	return &orderService{l: l, svc: svc}
}

func (h *orderService) Create(ctx context.Context, req *orderpb.CreateRequest) (*orderpb.CreateResponse, error) {
	items := make([]models.OrderItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = models.OrderItem{
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
		}
	}

	out, err := h.svc.CreateOrder(ctx, services.CreateOrderInput{
		UserID: req.UserId,
		Items:  items,
	})
	if err != nil {
		return nil, mapGRPCErrorCode(err)
	}

	return &orderpb.CreateResponse{
		OrderCode:  out.OrderCode,
		WorkflowId: out.WorkflowID,
		PaymentUrl: out.PaymentURL,
	}, nil
}

func (h *orderService) FindOne(ctx context.Context, req *orderpb.FindOneRequest) (*orderpb.FindOneResponse, error) {
	o, err := h.svc.FindOneOrder(ctx, services.GetOneOrderInput{
		ID:   req.GetId(),
		Code: req.GetCode(),
	})
	if err != nil {
		return nil, mapGRPCErrorCode(err)
	}

	return &orderpb.FindOneResponse{
		Order: toOrderDataProto(o),
	}, nil
}

func (h *orderService) FindMany(ctx context.Context, req *orderpb.FindManyRequest) (*orderpb.FindManyResponse, error) {
	out, err := h.svc.GetOrders(ctx, services.GetOrdersInput{
		UserID: req.GetUserId(),
		Status: stringFromOrderStatusProto(req.GetStatus()),
	})
	if err != nil {
		return nil, mapGRPCErrorCode(err)
	}

	os := make([]*orderpb.OrderData, len(out.Orders))
	for i, o := range out.Orders {
		os[i] = toOrderDataProto(o)
	}

	return &orderpb.FindManyResponse{
		Orders: os,
	}, nil
}

func (h *orderService) UpdateStatus(ctx context.Context, req *orderpb.UpdateStatusRequest) (*emptypb.Empty, error) {
	if err := h.svc.UpdateOrderStatus(ctx, services.UpdateOrderStatusInput{
		ID:     req.GetId(),
		Code:   req.GetCode(),
		Status: stringFromOrderStatusProto(req.GetStatus()),
	}); err != nil {
		return nil, mapGRPCErrorCode(err)
	}

	return &emptypb.Empty{}, nil
}

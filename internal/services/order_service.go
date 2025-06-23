package service

import (
	"context"
	"time"

	"github.com/vogiaan1904/order-svc/internal/models"
	repository "github.com/vogiaan1904/order-svc/internal/repositories"
	"github.com/vogiaan1904/order-svc/pkg/log"
	"github.com/vogiaan1904/order-svc/pkg/mongo"
	"github.com/vogiaan1904/order-svc/protogen/golang/product"

	"go.temporal.io/sdk/client"
)

const PrePaymentOrderTaskQueue = "PRE_PAYMENT_ORDER_TASK_QUEUE"

type implOrderService struct {
	l              log.Logger
	repo           repository.OrderRepository
	productSvc     product.ProductServiceClient
	temporalClient client.Client
}

func NewOrderService(l log.Logger, repo repository.OrderRepository, productSvc product.ProductServiceClient, temporalClient client.Client) OrderService {
	return &implOrderService{
		l:              l,
		repo:           repo,
		productSvc:     productSvc,
		temporalClient: temporalClient,
	}
}

func (svc *implOrderService) CreateOrder(ctx context.Context, input CreateOrderInput) (CreateOrderOutput, error) {
	pIDs := make([]string, len(input.Items))
	for _, item := range input.Items {
		pIDs = append(pIDs, item.ProductID)
	}

	resp, err := svc.productSvc.List(ctx, &product.ListRequest{Ids: pIDs})
	if err != nil {
		svc.l.Errorf(ctx, "orderSvc.Create.productSvc.List: %v", err)
		return CreateOrderOutput{}, err
	}

	if err := svc.validateOrderItems(ctx, input, resp); err != nil {
		svc.l.Warnf(ctx, "orderSvc.Create.validateOrderItems: %v", err)
		return CreateOrderOutput{}, err
	}

	oItems := make([]models.OrderItem, len(pIDs))
	total := 0.0
	for i, p := range resp.Products {
		iTotal := input.Items[i].Quantity * int32(p.Price)
		total += float64(iTotal)
		oItems[i] = models.OrderItem{
			ProductID:    p.Id,
			Quantity:     input.Items[i].Quantity,
			ProductPrice: p.Price,
			ProductName:  p.Name,
		}
	}

	code := svc.generateOrderCode()

	_, err = svc.repo.CreateOrder(ctx, repository.CreateOrderOptions{
		Code:        code,
		UserID:      input.UserID,
		Items:       oItems,
		TotalAmount: float64(total),
		Status:      models.OrderStatusPending,
	})
	if err != nil {
		svc.l.Errorf(ctx, "orderSvc.Create.repo.CreateOrder: %v", err)
		return CreateOrderOutput{}, err
	}

	wfID := "order_pre_payment_" + code
	wfParams := OrderWorkflowParams{
		OrderCode:   code,
		UserId:      input.UserID,
		TotalAmount: float64(total),
	}

	wfOpts := client.StartWorkflowOptions{
		ID:                       wfID,
		TaskQueue:                PrePaymentOrderTaskQueue,
		WorkflowExecutionTimeout: time.Hour * 24,
		WorkflowRunTimeout:       time.Hour * 24,
		WorkflowTaskTimeout:      time.Minute * 1,
	}

	we, err := svc.temporalClient.ExecuteWorkflow(ctx, wfOpts, "ProcessPrePaymentOrder", wfParams)
	if err != nil {
		svc.l.Errorf(ctx, "orderSvc.Create.temporalClient.ExecuteWorkflow: %v", err)
		return CreateOrderOutput{}, err
	}

	var paymentUrl string
	err = we.Get(ctx, &paymentUrl)
	if err != nil {
		svc.l.Errorf(ctx, "orderSvc.Create.temporalClient.ExecuteWorkflow.Get: %v", err)
		return CreateOrderOutput{}, err
	}

	return CreateOrderOutput{
		OrderCode:  code,
		WorkflowID: we.GetID(),
		PaymentURL: paymentUrl,
	}, nil
}

func (svc *implOrderService) FindOneOrder(ctx context.Context, input GetOneOrderInput) (models.Order, error) {
	o, err := svc.repo.FindOneOrder(ctx, repository.FindOneOrderOptions{
		ID: input.ID,
		GetOrdersFilter: repository.GetOrdersFilter{
			Code: input.Code,
		},
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			svc.l.Warnf(ctx, "orderSvc.FindOneOrder.repo.FindOneOrder: %v", err)
			return models.Order{}, ErrOrderNotFound
		}
		svc.l.Errorf(ctx, "orderSvc.FindOneOrder.repo.FindOneOrder: %v", err)
		return models.Order{}, err
	}

	return o, nil
}

func (svc *implOrderService) GetOrders(ctx context.Context, input GetOrdersInput) (GetOrdersOutput, error) {
	os, err := svc.repo.GetOrders(ctx, repository.GetOrdersOptions{
		GetOrdersFilter: repository.GetOrdersFilter{
			UserID: input.UserID,
			Status: input.Status,
		},
	})
	if err != nil {
		svc.l.Errorf(ctx, "OrderSvc.GetOrders.repo.GetOrders: %v", err)
		return GetOrdersOutput{}, err
	}

	return GetOrdersOutput{
		Orders: os,
	}, nil
}

func (svc *implOrderService) UpdateOrderStatus(ctx context.Context, input UpdateOrderStatusInput) error {
	o, err := svc.repo.FindOneOrder(ctx, repository.FindOneOrderOptions{
		GetOrdersFilter: repository.GetOrdersFilter{
			Code: input.Code,
		},
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			svc.l.Warnf(ctx, "OrderSvc.UpdateStatus.repo.FindOneOrder: %v", err)
			return ErrOrderNotFound
		}
		svc.l.Errorf(ctx, "OrderSvc.UpdateStatus.repo.FindOneOrder: %v", err)
		return err
	}

	o.Status = models.OrderStatus(input.Status)
	err = svc.repo.UpdateOrder(ctx, o)
	if err != nil {
		svc.l.Errorf(ctx, "OrderSvc.UpdateStatus.repo.UpdateOrder: %v", err)
		return err
	}

	return nil
}

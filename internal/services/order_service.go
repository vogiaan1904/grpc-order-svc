package service

import (
	"context"
	"time"

	repository "github.com/vogiaan1904/order-svc/internal/repositories"
	"github.com/vogiaan1904/order-svc/pkg/log"
	"github.com/vogiaan1904/order-svc/pkg/mongo"
	order "github.com/vogiaan1904/order-svc/protogen/golang/order"
	"github.com/vogiaan1904/order-svc/protogen/golang/product"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const OrderProcessingTaskQueue = "ORDER_PROCESSING_TASK_QUEUE"

type implOrderService struct {
	l              log.Logger
	repo           repository.OrderRepository
	productSvc     product.ProductServiceClient
	temporalClient client.Client
	order.UnimplementedOrderServiceServer
}

func NewOrderService(l log.Logger, repo repository.OrderRepository, productSvc product.ProductServiceClient, temporalClient client.Client) order.OrderServiceServer {
	return &implOrderService{
		l:              l,
		repo:           repo,
		productSvc:     productSvc,
		temporalClient: temporalClient,
	}
}

type OrderWorkflowParams struct {
	OrderID     string
	UserID      string
	Items       []*order.OrderItem
	TotalAmount float64
}

func (svc *implOrderService) Create(ctx context.Context, req *order.CreateRequest) (*order.CreateResponse, error) {
	svc.l.Infof(ctx, "Create order request received: UserID: %s, Items: %d", req.UserId, len(req.Items))
	pIDs := make([]string, len(req.Items))
	for _, item := range req.Items {
		pIDs = append(pIDs, item.ProductId)
	}

	resp, err := svc.productSvc.List(ctx, &product.ListRequest{Ids: pIDs})
	if err != nil {
		svc.l.Errorf(ctx, "Failed to get products: %v", err)
		return nil, err
	}

	if err := svc.validateOrderItems(ctx, req, resp); err != nil {
		svc.l.Warnf(ctx, "orderSvc.Create: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	oID := primitive.NewObjectID().Hex()
	workflowID := "order_" + oID

	oItems := make([]*order.OrderItem, len(pIDs))
	total := 0.0
	for i, p := range resp.Products {
		iTotal := req.Items[i].Quantity * int32(p.Price)
		total += float64(iTotal)
		oItems[i] = &order.OrderItem{
			ProductId:    p.Id,
			Quantity:     req.Items[i].Quantity,
			ProductPrice: p.Price,
			ProductName:  p.Name,
			TotalAmount:  float64(iTotal),
		}
	}

	workflowParams := OrderWorkflowParams{
		OrderID:     oID,
		UserID:      req.UserId,
		Items:       oItems,
		TotalAmount: float64(total),
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:                       workflowID,
		TaskQueue:                OrderProcessingTaskQueue,
		WorkflowExecutionTimeout: time.Hour * 24,
		WorkflowRunTimeout:       time.Hour * 24,
		WorkflowTaskTimeout:      time.Minute * 1,
	}

	svc.l.Infof(ctx, "Starting OrderProcessingWorkflow with ID: %s", workflowID)
	we, err := svc.temporalClient.ExecuteWorkflow(ctx, workflowOptions, "OrderProcessingWorkflow", workflowParams)
	if err != nil {
		svc.l.Errorf(ctx, "Failed to start OrderProcessingWorkflow: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to initiate order processing: %v", err)
	}

	svc.l.Infof(ctx, "OrderProcessingWorkflow started successfully. WorkflowID: %s, RunID: %s", we.GetID(), we.GetRunID())

	return &order.CreateResponse{
		OrderId:    oID,
		WorkflowId: we.GetID(),
	}, nil
}

func (svc *implOrderService) FindOne(ctx context.Context, req *order.FindOneRequest) (*order.FindOneResponse, error) {
	o, err := svc.repo.FindOneOrder(ctx, repository.FindOneOrderOptions{
		ID:         req.Id,
		FindFilter: repository.FindFilter{},
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			svc.l.Warnf(ctx, "OrderSvc.FindOne: %v", err)
			return nil, status.Errorf(codes.InvalidArgument, ErrOrderNotFound.Error())
		}
		svc.l.Errorf(ctx, "`Error FindOneOrder`: %v", err)
		return nil, err
	}

	var oItems []*order.OrderItem
	for _, item := range o.Items {
		oItems = append(oItems, &order.OrderItem{
			ProductId:    item.ProductID,
			Quantity:     item.Quantity,
			ProductPrice: item.ProductPrice,
			ProductName:  item.ProductName,
			TotalAmount:  item.ProductPrice * float64(item.Quantity),
		})
	}

	return &order.FindOneResponse{
		Order: &order.OrderData{
			Id:          o.ID.Hex(),
			UserId:      o.UserID,
			Items:       oItems,
			TotalAmount: o.TotalAmount,
			Status:      order.OrderStatus(order.OrderStatus_value[string(o.Status)]),
		},
	}, nil
}

func (svc *implOrderService) FindMany(ctx context.Context, req *order.FindManyRequest) (*order.FindManyResponse, error) {
	opt := svc.buildFindOptions(req)
	dbOrds, err := svc.repo.FindManyOrders(ctx, opt)
	if err != nil {
		svc.l.Errorf(ctx, "OrderSvc.FindMany: %v", err)
		return nil, err
	}
	svc.l.Debugf(ctx, "Found %d orders", len(dbOrds))

	ords := make([]*order.OrderData, len(dbOrds))
	for i, o := range dbOrds {
		ords[i] = svc.toOrderDataResp(o)
	}

	return &order.FindManyResponse{
		Orders: ords,
	}, nil
}

func (svc *implOrderService) UpdateStatus(ctx context.Context, req *order.UpdateStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}

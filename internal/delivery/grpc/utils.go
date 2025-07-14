package grpc

import (
	"github.com/vogiaan1904/order-svc/internal/models"
	orderpb "github.com/vogiaan1904/order-svc/protogen/golang/order"
)

func toOrderDataProto(o models.Order) *orderpb.OrderData {
	d := &orderpb.OrderData{
		Id:          o.ID.Hex(),
		UserId:      o.UserID,
		Items:       toOrderItemProto(o.Items),
		TotalAmount: o.TotalAmount,
		Status:      orderpb.OrderStatus(orderpb.OrderStatus_value[string(o.Status)]),
	}

	return d
}

func toOrderItemProto(items []models.OrderItem) []*orderpb.OrderItem {
	oItems := make([]*orderpb.OrderItem, len(items))
	for i, item := range items {
		oItems[i] = &orderpb.OrderItem{
			ProductId:    item.ProductID,
			Quantity:     item.Quantity,
			ProductPrice: item.ProductPrice,
		}
	}

	return oItems
}

func stringFromOrderStatusProto(status orderpb.OrderStatus) string {
	if status == orderpb.OrderStatus_ORDER_STATUS_UNSPECIFIED {
		return ""
	}

	str, ok := orderpb.OrderStatus_name[int32(status)]
	if !ok {
		return ""
	}

	return str
}

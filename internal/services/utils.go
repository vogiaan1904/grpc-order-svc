package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/vogiaan1904/order-svc/protogen/golang/product"
)

func (svc *implOrderService) validateOrderItems(ctx context.Context, input CreateOrderInput, resp *product.ListResponse) error {
	if len(resp.Products) != len(input.Items) {
		svc.l.Warnf(ctx, "orderSvc.validateOrderItems: %v", ErrProductNotFound)
		return ErrProductNotFound
	}

	pMap := make(map[string]*product.ProductData)
	for _, p := range resp.Products {
		pMap[p.Id] = p
	}

	for _, item := range input.Items {
		p, ok := pMap[item.ProductID]
		if !ok {
			svc.l.Warnf(ctx, "orderSvc.validateOrderItems: %v", ErrProductNotFound)
			return ErrProductNotFound
		}
		if p.TotalStock < int32(item.Quantity) {
			svc.l.Warnf(ctx, "orderSvc.validateOrderItems: Product %v does not have enough stock. Requested: %d, Available: %d", item.ProductID, item.Quantity, p.TotalStock)
			return ErrInsufficientStock
		}
	}

	return nil
}

func (svc *implOrderService) generateOrderCode() string {
	now := time.Now()
	dateStr := now.Format("02012006")
	randomNum := rand.Intn(100000000)
	randomStr := fmt.Sprintf("%08d", randomNum)

	return dateStr + randomStr
}

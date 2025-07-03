package http

import (
	"github.com/vogiaan1904/order-svc/internal/middlewares"
	"github.com/vogiaan1904/order-svc/internal/services"
	pkgLog "github.com/vogiaan1904/order-svc/pkg/log"
)

type orderHandler struct {
	l   pkgLog.Logger
	svc services.OrderService
	mw  middlewares.MiddlewareManager
}

func NewOrderHandler(l pkgLog.Logger, svc services.OrderService, mw middlewares.MiddlewareManager) *orderHandler {
	return &orderHandler{l: l, svc: svc, mw: mw}
}

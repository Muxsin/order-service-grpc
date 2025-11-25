package handler

import (
	"context"

	orders "github.com/muxsin/kitchen/services/common/genproto/orders/protobuf"
	"github.com/muxsin/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrderGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewOrderGrpcHandler(grpc *grpc.Server, ordersService types.OrderService) {
	grpcHandler := &OrderGrpcHandler{
		orderService: ordersService,
	}

	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrderGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
	o := h.orderService.GetOrders(ctx)
	return &orders.GetOrderResponse{
		Orders: o,
	}, nil
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    1,
		CustomerID: 1,
		ProductID:  1,
		Quantity:   1,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}

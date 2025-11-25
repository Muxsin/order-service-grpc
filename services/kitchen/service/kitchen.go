package service

import (
	"context"

	orders "github.com/muxsin/kitchen/services/common/genproto/orders/protobuf"
	"google.golang.org/grpc"
)

type KitchenService struct {
	client orders.OrderServiceClient
}

func NewKitchenService(conn *grpc.ClientConn) *KitchenService {
	return &KitchenService{
		client: orders.NewOrderServiceClient(conn),
	}
}

func (s *KitchenService) CreateOrder(ctx context.Context, order *orders.CreateOrderRequest) error {
	_, err := s.client.CreateOrder(ctx, order)

	return err
}

func (s *KitchenService) GetOrders(ctx context.Context, order *orders.GetOrdersRequest) ([]*orders.Order, error) {
	res, err := s.client.GetOrders(ctx, order)
	if err != nil {
		return nil, err
	}

	return res.GetOrders(), nil
}

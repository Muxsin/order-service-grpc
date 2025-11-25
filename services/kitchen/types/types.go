package types

import (
	"context"

	orders "github.com/muxsin/kitchen/services/common/genproto/orders/protobuf"
)

type KitchenService interface {
	CreateOrder(context.Context, *orders.CreateOrderRequest) error
	GetOrders(context.Context, *orders.GetOrdersRequest) ([]*orders.Order, error)
}

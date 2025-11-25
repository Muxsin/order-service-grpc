package types

import (
	"context"

	orders "github.com/muxsin/kitchen/services/common/genproto/orders/protobuf"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}

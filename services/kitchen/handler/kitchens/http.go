package handler

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	orders "github.com/muxsin/kitchen/services/common/genproto/orders/protobuf"
	"github.com/muxsin/kitchen/services/kitchen/types"
)

type KitchenHttpHandler struct {
	kitchenService types.KitchenService
}

func NewKitchenHttpHandler(kitchenService types.KitchenService) *KitchenHttpHandler {
	return &KitchenHttpHandler{
		kitchenService: kitchenService,
	}
}

func (h *KitchenHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("GET /", h.GetOrders)
}

func (h *KitchenHttpHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	orderReq := &orders.CreateOrderRequest{
		CustomerID: 1,
		ProductID:  1,
		Quantity:   2,
	}

	err := h.kitchenService.CreateOrder(ctx, orderReq)
	if err != nil {
		log.Fatalf("client error: %v", err)
	}

	orderRes := &orders.GetOrdersRequest{
		CustomerID: 2,
	}

	res, err := h.kitchenService.GetOrders(r.Context(), orderRes)
	if err != nil {
		log.Fatalf("client error: %v", err)
	}

	t := template.Must(template.New("orders").Parse(ordersTemplate))

	if err := t.Execute(w, res); err != nil {
		log.Fatalf("template error: %v", err)
	}
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`

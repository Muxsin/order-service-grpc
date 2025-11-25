package handler

import (
	"encoding/json"
	"net/http"

	orders "github.com/muxsin/kitchen/services/common/genproto/orders/protobuf"
	"github.com/muxsin/kitchen/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewOrdersHttpHandler(orderService types.OrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{
		ordersService: orderService,
	}
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	order := &orders.Order{
		OrderID:    1,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders.CreateOrderResponse{
		Status: "success",
	})
}

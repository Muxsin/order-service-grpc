package main

import (
	"log"
	"net/http"

	handler "github.com/muxsin/kitchen/services/orders/handler/orders"
	"github.com/muxsin/kitchen/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (hs *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewOrdersHttpHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting server on", hs.addr)

	return http.ListenAndServe(hs.addr, router)
}

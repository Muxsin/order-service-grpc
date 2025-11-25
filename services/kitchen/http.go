package main

import (
	"log"
	"net/http"

	handler "github.com/muxsin/kitchen/services/kitchen/handler/kitchens"
	"github.com/muxsin/kitchen/services/kitchen/service"
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

	conn := NewGrpcClient(":9000")
	defer conn.Close()

	kitchenService := service.NewKitchenService(conn)
	kitchenHandler := handler.NewKitchenHttpHandler(kitchenService)
	kitchenHandler.RegisterRouter(router)

	log.Println("Starting server on", hs.addr)

	return http.ListenAndServe(hs.addr, router)
}

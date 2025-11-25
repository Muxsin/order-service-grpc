package main

import (
	"log"
	"net"

	handler "github.com/muxsin/kitchen/services/orders/handler/orders"
	"github.com/muxsin/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (gs *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", gs.addr)
	if err != nil {
		log.Printf("Failed to listen: %v", err)

		return err
	}

	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewOrderGrpcHandler(grpcServer, orderService)

	log.Print("Starting gRPC server on", gs.addr)

	return grpcServer.Serve(lis)
}

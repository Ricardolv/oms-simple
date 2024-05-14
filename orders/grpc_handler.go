package main

import (
	"context"
	pb "github.com/Ricardolv/commons/api"
	"google.golang.org/grpc"
	"log"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGrpcHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(c context.Context, p *pb.CreateOderRequest) (*pb.Order, error) {

	log.Println("New order received! Order %v", p)

	o := &pb.Order{
		ID: "42",
	}

	return o, nil
}

package main

import (
	"context"
	"github.com/Ricardolv/commons"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	grpcAddr = commons.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()
	listen, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	defer listen.Close()

	store := NewStore()
	svc := NewService(store)
	NewGrpcHandler(grpcServer)

	svc.CreateOrder(context.Background())

	log.Println("GRPC Server Started at", grpcAddr)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf(err.Error())
	}

}

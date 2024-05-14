package main

import (
	"context"
	pb "github.com/Ricardolv/commons/api"
)

type OrdersService interface {
	CreateOrder(ctx context.Context) error
	ValidateOrder(context.Context, *pb.CreateOderRequest) error
}

type OrdersStore interface {
	Create(ctx context.Context) error
}

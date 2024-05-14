package main

import (
	"context"
	"github.com/Ricardolv/commons"
	pb "github.com/Ricardolv/commons/api"
	"log"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}

func (s *service) ValidateOrder(context context.Context, request *pb.CreateOderRequest) error {
	if len(request.Items) == 0 {
		return commons.ErrNoItems
	}

	mergedItems := mergeItemsQuantities(request.Items)
	log.Print(mergedItems)

	// validate with the stock service

	return nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQuatity) []*pb.ItemsWithQuatity {
	merged := make([]*pb.ItemsWithQuatity, 0)

	for _, item := range items {

		found := false
		for _, finalItem := range merged {
			if finalItem.ID == item.ID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}

		if !found {
			merged = append(merged, item)
		}
	}

	return merged
}

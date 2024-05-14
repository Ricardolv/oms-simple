package main

import (
	"errors"
	"github.com/Ricardolv/commons"
	pb "github.com/Ricardolv/commons/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandlerCretaeOrders)
}

func (h *handler) HandlerCretaeOrders(resp http.ResponseWriter, req *http.Request) {

	customerID := req.PathValue("customerID")

	var items []*pb.ItemsWithQuatity
	if err := commons.ReadJSON(req, &items); err != nil {
		commons.WriterError(resp, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateItems(items); err != nil {
		commons.WriterError(resp, http.StatusBadRequest, err.Error())
	}

	order, err := h.client.CreateOrder(req.Context(), &pb.CreateOderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	rStatus := status.Convert(err)

	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			commons.WriterError(resp, http.StatusBadRequest, rStatus.Message())
			return
		}

		commons.WriterError(resp, http.StatusInternalServerError, err.Error())
		return
	}

	commons.WriteJSON(resp, http.StatusOK, order)
}

func validateItems(items []*pb.ItemsWithQuatity) error {

	if len(items) == 0 {
		return commons.ErrNoItems
	}

	for _, item := range items {
		if item.ID == "" {
			return errors.New("item ID is required")
		}

		if item.Quantity <= 0 {
			return errors.New("item must have a valid quantity")
		}
	}

	return nil
}

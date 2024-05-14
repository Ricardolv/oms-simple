package main

import (
	"github.com/Ricardolv/commons"
	pb "github.com/Ricardolv/commons/api"
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

	h.client.CreateOrder(req.Context(), &pb.CreateOderRequest{
		CustomerID: customerID,
		Items:      items,
	})
}

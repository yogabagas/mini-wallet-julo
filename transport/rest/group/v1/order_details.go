package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogabagas/jatis-BE/transport/rest/handler"
)

func NewOrderDetailsGroupV1(v1 *mux.Router, h handler.HandlerImpl) {
	v1.HandleFunc("/order-details", h.GetOrderDetails).Methods(http.MethodGet)
	v1.HandleFunc("/order-details", h.CreateOrderDetails).Methods(http.MethodPost)
}

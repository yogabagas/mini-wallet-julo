package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogabagas/jatis-BE/transport/rest/handler"
)

func NewShippingMethodsGroupV1(v1 *mux.Router, h handler.HandlerImpl) {
	v1.HandleFunc("/shipping-methods", h.CreateShippingMethods).Methods(http.MethodPost)
}

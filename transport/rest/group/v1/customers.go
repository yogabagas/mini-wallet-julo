package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogabagas/jatis-BE/transport/rest/handler"
)

func NewCustomersGroupV1(v1 *mux.Router, h handler.HandlerImpl) {
	v1.HandleFunc("/customers", h.CreateCustomers).Methods(http.MethodPost)
}

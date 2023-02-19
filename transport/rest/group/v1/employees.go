package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yogabagas/jatis-BE/transport/rest/handler"
)

func NewEmployeesGroupV1(v1 *mux.Router, h handler.HandlerImpl) {
	v1.HandleFunc("/employees", h.CreateEmployees).Methods(http.MethodPost)
}

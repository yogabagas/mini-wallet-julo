package rest

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/yogabagas/jatis-BE/registry"
	groupV1 "github.com/yogabagas/jatis-BE/transport/rest/group/v1"
	"github.com/yogabagas/jatis-BE/transport/rest/handler"
)

type Options struct {
	Port         string
	Db           *sql.DB
	Address      string
	Mux          *mux.Router
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

type Handler struct {
	options *Options
}

func NewRest(o *Options) *Handler {
	reg := registry.NewRegistry(
		registry.NewSQLConn(o.Db),
	)
	appController := reg.NewAppController()

	handlerImpl := handler.HandlerImpl{
		Controller: appController,
	}

	r := mux.NewRouter()

	r.Path("/health").HandlerFunc(handlerImpl.HealthCheck)

	v1 := r.PathPrefix("/v1").Subrouter()

	groupV1.NewCustomersGroupV1(v1, handlerImpl)
	groupV1.NewOrderDetailsGroupV1(v1, handlerImpl)
	groupV1.NewEmployeesGroupV1(v1, handlerImpl)
	groupV1.NewProductsGroupV1(v1, handlerImpl)
	groupV1.NewShippingMethodsGroupV1(v1, handlerImpl)
	groupV1.NewOrdersGroupV1(v1, handlerImpl)

	o.Mux = r

	return &Handler{options: o}
}

func (h *Handler) Serve() {
	log.Printf("server serve at port %s", h.options.Port)
	http.ListenAndServe(h.options.Port, h.options.Mux)
}

package rest

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/yogabagas/mini-wallet-julo/registry"
	groupV1 "github.com/yogabagas/mini-wallet-julo/transport/rest/group/v1"
	"github.com/yogabagas/mini-wallet-julo/transport/rest/handler"
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
	middleware := reg.NewMiddleware()

	handlerImpl := handler.HandlerImpl{
		Controller: appController,
		Middleware: middleware,
	}

	r := mux.NewRouter()

	r.Path("/health").HandlerFunc(handlerImpl.HealthCheck)

	apiPath := r.PathPrefix("/api")

	v1 := apiPath.PathPrefix("/v1").Subrouter()

	groupV1.NewWalletsGroupV1(v1, handlerImpl, middleware)

	o.Mux = r

	return &Handler{options: o}
}

func (h *Handler) Serve() {
	log.Printf("server serve at port %s", h.options.Port)
	http.ListenAndServe(h.options.Port, h.options.Mux)
}

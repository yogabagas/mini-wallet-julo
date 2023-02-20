package handler

import (
	"github.com/yogabagas/mini-wallet-julo/adapter/controller"
	"github.com/yogabagas/mini-wallet-julo/transport/rest/middleware"
)

type HandlerImpl struct {
	Controller controller.AppController
	Middleware middleware.Middleware
}

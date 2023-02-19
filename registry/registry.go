package registry

import (
	"database/sql"

	"github.com/yogabagas/jatis-BE/adapter/controller"
)

type module struct {
	sql *sql.DB
}

type Controller interface {
	NewAppController() controller.AppController
}

type Option func(*module)

func NewRegistry(options ...Option) Controller {
	m := &module{}
	for _, o := range options {
		o(m)
	}
	return m
}

func (m *module) NewAppController() controller.AppController {
	return controller.AppController{
		CustomersController:       m.NewCustomersController(),
		OrderDetailsController:    m.NewOrderDetailsController(),
		EmployeesController:       m.NewEmployeesController(),
		ProductsController:        m.NewProductsController(),
		ShippingMethodsController: m.NewShippingMethodsController(),
		OrdersController:          m.NewOrdersController(),
		WalletsController:         m.NewWalletsController(),
	}
}

func NewSQLConn(mdb *sql.DB) Option {
	return func(m *module) {
		m.sql = mdb
	}
}

package registry

import (
	"database/sql"

	"github.com/yogabagas/mini-wallet-julo/adapter/controller"
	"github.com/yogabagas/mini-wallet-julo/domain/repository"
	"github.com/yogabagas/mini-wallet-julo/transport/rest/middleware"
)

type module struct {
	sql *sql.DB
}

type Controller interface {
	NewAppController() controller.AppController
	NewMiddleware() middleware.Middleware
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
		WalletsController:                m.NewWalletsController(),
		WalletBalanceHistoriesController: m.NewWalletBalanceHistoriesController(),
	}
}

func (m *module) NewMiddleware() middleware.Middleware {
	return middleware.NewMiddleware()
}

func (m *module) NewRepositoryRegistry() repository.RepositoryRegistry {
	return repository.NewRepositoryRegistry(m.sql)
}

func NewSQLConn(mdb *sql.DB) Option {
	return func(m *module) {
		m.sql = mdb
	}
}

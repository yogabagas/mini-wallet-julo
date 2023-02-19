package registry

import (
	"github.com/yogabagas/jatis-BE/adapter/controller"
	"github.com/yogabagas/jatis-BE/domain/repository"
	orders "github.com/yogabagas/jatis-BE/service/orders/usecase"
)

func (m *module) NewOrdersRepository() repository.OrdersRepository {
	return repository.NewOrdersRepository(m.sql)
}

func (m *module) NewOrdersUsecase() orders.OrdersUsecase {
	return orders.NewOrdersUsecase(m.NewOrdersRepository())
}

func (m *module) NewOrdersController() controller.OrdersController {
	return controller.NewOrdersController(m.NewOrdersUsecase())
}

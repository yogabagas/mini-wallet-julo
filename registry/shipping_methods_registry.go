package registry

import (
	"github.com/yogabagas/jatis-BE/adapter/controller"
	"github.com/yogabagas/jatis-BE/domain/repository"
	shippingMethods "github.com/yogabagas/jatis-BE/service/shippingmethods/usecase"
)

func (m *module) NewShippingMethodsRepository() repository.ShippingMethodsRepository {
	return repository.NewShippingMethodsRepository(m.sql)
}

func (m *module) NewShippingMethodsUsecase() shippingMethods.ShippingMethodsUsecase {
	return shippingMethods.NewShippingMethodsUsecase(m.NewShippingMethodsRepository())
}

func (m *module) NewShippingMethodsController() controller.ShippingMethodsController {
	return controller.NewShippingMethodsController(m.NewShippingMethodsUsecase())
}

package registry

import (
	"github.com/yogabagas/jatis-BE/adapter/controller"
	"github.com/yogabagas/jatis-BE/domain/repository"
	customers "github.com/yogabagas/jatis-BE/service/customers/usecase"
)

func (m *module) NewCustomersRepository() repository.CustomersRepository {
	return repository.NewCustomersRepository(m.sql)
}

func (m *module) NewCustomersUsecase() customers.CustomersUsecase {
	return customers.NewCustomersUsecase(m.NewCustomersRepository())
}

func (m *module) NewCustomersController() controller.CustomersController {
	return controller.NewCustomersController(m.NewCustomersUsecase())
}

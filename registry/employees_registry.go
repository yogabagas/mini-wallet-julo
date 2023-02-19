package registry

import (
	"github.com/yogabagas/jatis-BE/adapter/controller"
	"github.com/yogabagas/jatis-BE/domain/repository"
	employees "github.com/yogabagas/jatis-BE/service/employees/usecase"
)

func (m *module) NewEmployeesRepository() repository.EmployeesRepository {
	return repository.NewEmployeesRepository(m.sql)
}

func (m *module) NewEmployeesUsecase() employees.EmployeesUsecase {
	return employees.NewEmployeesUsecase(m.NewEmployeesRepository())
}

func (m *module) NewEmployeesController() controller.EmployeesController {
	return controller.NewEmployeesController(m.NewEmployeesUsecase())
}

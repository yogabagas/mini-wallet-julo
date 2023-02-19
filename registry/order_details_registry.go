package registry

import (
	"github.com/yogabagas/jatis-BE/adapter/controller"
	"github.com/yogabagas/jatis-BE/domain/repository"
	orderDetails "github.com/yogabagas/jatis-BE/service/orderdetails/usecase"
)

func (m *module) NewOrderDetailsRepository() repository.OrderDetaisRepository {
	return repository.NewOrderDetailsRepository(m.sql)
}

func (m *module) NewOrderDetailsUsecase() orderDetails.OrderDetailsUsecase {
	return orderDetails.NewOrderDetailsUsecase(m.NewOrderDetailsRepository())
}

func (m *module) NewOrderDetailsController() controller.OrderDetailsController {
	return controller.NewOrderDetailsController(m.NewOrderDetailsUsecase())
}

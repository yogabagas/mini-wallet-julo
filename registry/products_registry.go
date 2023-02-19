package registry

import (
	"github.com/yogabagas/jatis-BE/adapter/controller"
	"github.com/yogabagas/jatis-BE/domain/repository"
	products "github.com/yogabagas/jatis-BE/service/products/usecase"
)

func (m *module) NewProductsRepository() repository.ProductsRepository {
	return repository.NewProductsRepository(m.sql)
}

func (m *module) NewProductsUsecase() products.ProductsUsecase {
	return products.NewProductsUsecase(m.NewProductsRepository())
}

func (m *module) NewProductsController() controller.ProductsController {
	return controller.NewProductsController(m.NewProductsUsecase())
}

package registry

import (
	"github.com/yogabagas/jatis-BE/adapter/controller"
	"github.com/yogabagas/jatis-BE/domain/repository"
	"github.com/yogabagas/jatis-BE/service/wallets/presenter"
	wallets "github.com/yogabagas/jatis-BE/service/wallets/usecase"
)

func (m *module) NewRepositoryRegistry() repository.RepositoryRegistry {
	return repository.NewRepositoryRegistry(m.sql)
}

func (m *module) NewWalletsPresenter() presenter.WalletsPresenter {
	return presenter.NewWalletsPresenter()
}

func (m *module) NewWalletsUsecase() wallets.WalletsUsecase {
	return wallets.NewWalletsUsecase(m.NewRepositoryRegistry(), m.NewWalletsPresenter())
}

func (m *module) NewWalletsController() controller.WalletsController {
	return controller.NewWalletsController(m.NewWalletsUsecase())
}

package registry

import (
	"github.com/yogabagas/mini-wallet-julo/adapter/controller"
	"github.com/yogabagas/mini-wallet-julo/service/wallets/presenter"
	wallets "github.com/yogabagas/mini-wallet-julo/service/wallets/usecase"
)

func (m *module) NewWalletsPresenter() presenter.WalletsPresenter {
	return presenter.NewWalletsPresenter()
}

func (m *module) NewWalletsUsecase() wallets.WalletsUsecase {
	return wallets.NewWalletsUsecase(m.NewRepositoryRegistry(), m.NewWalletsPresenter())
}

func (m *module) NewWalletsController() controller.WalletsController {
	return controller.NewWalletsController(m.NewWalletsUsecase())
}

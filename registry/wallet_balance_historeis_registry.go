package registry

import (
	"github.com/yogabagas/mini-wallet-julo/adapter/controller"
	"github.com/yogabagas/mini-wallet-julo/service/wallethistories/presenter"
	walletBalanceHistories "github.com/yogabagas/mini-wallet-julo/service/wallethistories/usecase"
)

func (m *module) NewWalletBalanceHistoriesPresenter() presenter.WalletBalanceHistoriesPresenter {
	return presenter.NewWalletBalanceHistoriesPresenter()
}

func (m *module) NewWalletBalanceHistoriesUsecase() walletBalanceHistories.WalletBalanceHistoriesUsecase {
	return walletBalanceHistories.NewWalletBalanceHistories(m.NewRepositoryRegistry(), m.NewWalletBalanceHistoriesPresenter())
}

func (m *module) NewWalletBalanceHistoriesController() controller.WalletBalanceHistoriesController {
	return controller.NewWalletBalanceHistoriesController(m.NewWalletBalanceHistoriesUsecase())
}

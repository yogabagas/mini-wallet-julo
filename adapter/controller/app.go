package controller

type AppController struct {
	WalletsController                interface{ WalletsController }
	WalletBalanceHistoriesController interface {
		WalletBalanceHistoriesController
	}
}

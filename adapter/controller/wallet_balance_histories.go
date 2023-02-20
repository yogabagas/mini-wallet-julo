package controller

import (
	"context"

	"github.com/yogabagas/mini-wallet-julo/domain/service"
	walletBalanceHistories "github.com/yogabagas/mini-wallet-julo/service/wallethistories/usecase"
)

type WalletBalanceHistoriesControllerImpl struct {
	walletBalanceHistoriesUsecase walletBalanceHistories.WalletBalanceHistoriesUsecase
}

type WalletBalanceHistoriesController interface {
	Deposit(ctx context.Context, req service.WalletDepositRequest) (resp service.WalletDepositResponse, err error)
	Withdrawal(ctx context.Context, req service.WalletWithdrawalsRequest) (resp service.WalletWithdrawalsResponse, err error)
	Transactions(ctx context.Context, req service.WalletTransactionsRequest) (resp service.WalletTransactionsResponse, err error)
}

func NewWalletBalanceHistoriesController(walletBalanceHistoriesUsecase walletBalanceHistories.WalletBalanceHistoriesUsecase) WalletBalanceHistoriesController {
	return &WalletBalanceHistoriesControllerImpl{walletBalanceHistoriesUsecase: walletBalanceHistoriesUsecase}
}

func (wc *WalletBalanceHistoriesControllerImpl) Deposit(ctx context.Context, req service.WalletDepositRequest) (resp service.WalletDepositResponse, err error) {
	return wc.walletBalanceHistoriesUsecase.Deposit(ctx, req)
}

func (wc *WalletBalanceHistoriesControllerImpl) Withdrawal(ctx context.Context, req service.WalletWithdrawalsRequest) (resp service.WalletWithdrawalsResponse, err error) {
	return wc.walletBalanceHistoriesUsecase.Withdrawal(ctx, req)
}

func (wc *WalletBalanceHistoriesControllerImpl) Transactions(ctx context.Context, req service.WalletTransactionsRequest) (resp service.WalletTransactionsResponse, err error) {
	return wc.walletBalanceHistoriesUsecase.Transactions(ctx, req)
}

package usecase

import (
	"context"
	"errors"

	"github.com/yogabagas/mini-wallet-julo/domain/model"
	"github.com/yogabagas/mini-wallet-julo/domain/repository"
	"github.com/yogabagas/mini-wallet-julo/domain/service"
	"github.com/yogabagas/mini-wallet-julo/service/wallethistories/presenter"
	"github.com/yogabagas/mini-wallet-julo/shared/constant"
)

type WalletBalanceHistoriesUsecaseImpl struct {
	repoRegistry repository.RepositoryRegistry
	presenter    presenter.WalletBalanceHistoriesPresenter
}

type WalletBalanceHistoriesUsecase interface {
	Deposit(ctx context.Context, req service.WalletDepositRequest) (resp service.WalletDepositResponse, err error)
	Withdrawal(ctx context.Context, req service.WalletWithdrawalsRequest) (resp service.WalletWithdrawalsResponse, err error)
	Transactions(ctx context.Context, req service.WalletTransactionsRequest) (resp service.WalletTransactionsResponse, err error)
}

func NewWalletBalanceHistories(repoRegistry repository.RepositoryRegistry, presenter presenter.WalletBalanceHistoriesPresenter) WalletBalanceHistoriesUsecase {
	return &WalletBalanceHistoriesUsecaseImpl{repoRegistry: repoRegistry, presenter: presenter}
}

func (wu *WalletBalanceHistoriesUsecaseImpl) Deposit(ctx context.Context, req service.WalletDepositRequest) (resp service.WalletDepositResponse, err error) {
	walletsRepo := wu.repoRegistry.GetWalletsRepository()

	walletResp, err := walletsRepo.ReadWalletByToken(ctx, model.ReadWalletByTokenRequest{Token: req.Token})
	if err != nil {
		return resp, err
	}

	if walletResp.Status == constant.Disabled.Int() {
		return resp, errors.New("400 - wallet disabled")
	}

	var InTransaction = func(r repository.RepositoryRegistry) (out interface{}, err error) {

		walletBalanceHistoriesRepo := wu.repoRegistry.GetWalletBalanceHistoriesRepository()

		reqWalletBalanceHistories := model.InsertWalletBalanceHistoriesRequest{
			WalletID:    walletResp.ID,
			ReferenceID: req.ReferenceID,
			Amount:      float64(req.Amount),
			Type:        constant.Deposit.Int(),
			CreatedBy:   walletResp.OwnedBy,
			UpdatedBy:   walletResp.OwnedBy,
		}

		err = walletBalanceHistoriesRepo.InsertWalletBalanceHistories(ctx, reqWalletBalanceHistories)
		if err != nil {
			return resp, err
		}

		reqWallet := model.UpdateAmountWalletByIDRequest{
			ID: walletResp.ID,
		}

		err = walletsRepo.UpdateAmountWalletByID(ctx, reqWallet)
		if err != nil {
			return resp, err
		}

		return nil, nil
	}

	_, err = wu.repoRegistry.DoInTransaction(ctx, InTransaction)
	if err != nil {
		return resp, err
	}

	return wu.presenter.Deposit(ctx, float64(req.Amount), req.ReferenceID, walletResp)
}

func (wu *WalletBalanceHistoriesUsecaseImpl) Withdrawal(ctx context.Context, req service.WalletWithdrawalsRequest) (resp service.WalletWithdrawalsResponse, err error) {
	walletsRepo := wu.repoRegistry.GetWalletsRepository()

	walletResp, err := walletsRepo.ReadWalletByToken(ctx, model.ReadWalletByTokenRequest{Token: req.Token})
	if err != nil {
		return resp, err
	}

	if walletResp.Status == constant.Disabled.Int() {
		return resp, errors.New("400 - wallet disabled")
	}

	if (walletResp.Balance + float64(req.Amount)) < 0 {
		return resp, errors.New("400 - insufficient balance")
	}

	var InTransaction = func(r repository.RepositoryRegistry) (out interface{}, err error) {

		walletBalanceHistoriesRepo := wu.repoRegistry.GetWalletBalanceHistoriesRepository()

		reqWalletBalanceHistories := model.InsertWalletBalanceHistoriesRequest{
			WalletID:    walletResp.ID,
			ReferenceID: req.ReferenceID,
			Amount:      float64(req.Amount),
			Type:        constant.Withdrawal.Int(),
			CreatedBy:   walletResp.OwnedBy,
			UpdatedBy:   walletResp.OwnedBy,
		}

		err = walletBalanceHistoriesRepo.InsertWalletBalanceHistories(ctx, reqWalletBalanceHistories)
		if err != nil {
			return resp, err
		}

		reqWallet := model.UpdateAmountWalletByIDRequest{
			ID: walletResp.ID,
		}

		err = walletsRepo.UpdateAmountWalletByID(ctx, reqWallet)
		if err != nil {
			return resp, err
		}

		return nil, nil
	}

	_, err = wu.repoRegistry.DoInTransaction(ctx, InTransaction)
	if err != nil {
		return resp, err
	}

	return wu.presenter.Withdrawal(ctx, float64(req.Amount), req.ReferenceID, walletResp)
}

func (wu *WalletBalanceHistoriesUsecaseImpl) Transactions(ctx context.Context, req service.WalletTransactionsRequest) (resp service.WalletTransactionsResponse, err error) {

	walletRepository := wu.repoRegistry.GetWalletsRepository()

	walletResp, err := walletRepository.ReadWalletByToken(ctx, model.ReadWalletByTokenRequest{Token: req.Token})
	if err != nil {
		return resp, err
	}

	walletBalanceHistoriesRepo := wu.repoRegistry.GetWalletBalanceHistoriesRepository()

	historiesResp, err := walletBalanceHistoriesRepo.ReadWalletBalanceHisoriesByWalletID(ctx, model.ReadWalletBalanceHisoriesByWalletIDRequest{
		WalletID: walletResp.ID,
	})
	if err != nil {
		return resp, err
	}

	return wu.presenter.Transactions(ctx, historiesResp)
}

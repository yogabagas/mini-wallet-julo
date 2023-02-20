package usecase

import (
	"context"
	"time"

	"github.com/yogabagas/mini-wallet-julo/domain/model"
	"github.com/yogabagas/mini-wallet-julo/domain/repository"
	"github.com/yogabagas/mini-wallet-julo/domain/service"
	"github.com/yogabagas/mini-wallet-julo/service/wallets/presenter"
	"github.com/yogabagas/mini-wallet-julo/shared/constant"
)

type WalletsUsecaseImpl struct {
	repoRegistry repository.RepositoryRegistry
	presenter    presenter.WalletsPresenter
}

type WalletsUsecase interface {
	InitWallets(ctx context.Context, req service.InitWalletRequest) (resp service.InitWalletResponse, err error)
	EnabledWallets(ctx context.Context, req service.EnableWalletRequest) (resp service.EnableWalletResponse, err error)
	ViewWalletsBalance(ctx context.Context, req service.ViewWalletBalanceRequest) (resp service.ViewWalletBalanceResponse, err error)
	DisabledWallets(ctx context.Context, req service.DisableWalletRequest) (resp service.DisableWalletResponse, err error)
}

func NewWalletsUsecase(repoRegistry repository.RepositoryRegistry, presenter presenter.WalletsPresenter) WalletsUsecase {
	return &WalletsUsecaseImpl{
		repoRegistry: repoRegistry,
		presenter:    presenter}
}

func (wu *WalletsUsecaseImpl) InitWallets(ctx context.Context, req service.InitWalletRequest) (resp service.InitWalletResponse, err error) {

	walletRepository := wu.repoRegistry.GetWalletsRepository()

	reqWallet := model.CreateWalletRequest{
		ID:        req.ID,
		OwnedBy:   req.CustomerID,
		Token:     req.Token,
		Status:    constant.Disabled.Int(),
		Balance:   0,
		CreatedBy: req.CustomerID,
		UpdatedBy: req.CustomerID,
	}

	err = walletRepository.CreateWallet(ctx, reqWallet)
	if err != nil {
		return resp, err
	}

	return wu.presenter.InitWallets(ctx, reqWallet)
}

func (wu *WalletsUsecaseImpl) EnabledWallets(ctx context.Context, req service.EnableWalletRequest) (resp service.EnableWalletResponse, err error) {

	walletRepository := wu.repoRegistry.GetWalletsRepository()

	respWallets, err := walletRepository.ReadWalletByToken(ctx, model.ReadWalletByTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return resp, err
	}

	now := time.Now()

	if respWallets.ID != "" {
		reqWallet := model.UpdateWalletByTokenRequest{
			Status:    constant.Enabled.Int(),
			EnabledAt: now,
			UpdatedBy: respWallets.OwnedBy,
			Token:     req.Token,
		}

		err = walletRepository.UpdateWalletByToken(ctx, reqWallet)
		if err != nil {
			return resp, err
		}
	}

	respWallets.Status = constant.Enabled.Int()
	respWallets.EnabledAt = now

	return wu.presenter.EnabledWallets(ctx, respWallets)
}

func (wu *WalletsUsecaseImpl) ViewWalletsBalance(ctx context.Context, req service.ViewWalletBalanceRequest) (resp service.ViewWalletBalanceResponse, err error) {

	walletRepository := wu.repoRegistry.GetWalletsRepository()

	reqWallet := model.ReadWalletByTokenRequest{
		Token: req.Token,
	}

	respWallet, err := walletRepository.ReadWalletByToken(ctx, reqWallet)
	if err != nil {
		return resp, err
	}

	return wu.presenter.ViewWalletsBalance(ctx, respWallet)
}

func (wu *WalletsUsecaseImpl) DisabledWallets(ctx context.Context, req service.DisableWalletRequest) (resp service.DisableWalletResponse, err error) {
	walletRepository := wu.repoRegistry.GetWalletsRepository()

	respWallets, err := walletRepository.ReadWalletByToken(ctx, model.ReadWalletByTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return resp, err
	}

	if respWallets.ID != "" {
		reqWallet := model.UpdateWalletByTokenRequest{
			Status:    constant.Disabled.Int(),
			UpdatedBy: respWallets.OwnedBy,
			Token:     req.Token,
		}

		err = walletRepository.UpdateWalletByToken(ctx, reqWallet)
		if err != nil {
			return resp, err
		}
	}

	respWallets.Status = constant.Disabled.Int()

	return wu.presenter.DisabledWallets(ctx, respWallets)

}

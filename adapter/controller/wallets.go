package controller

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/yogabagas/mini-wallet-julo/domain/service"
	wallets "github.com/yogabagas/mini-wallet-julo/service/wallets/usecase"
)

type WalletsControllerImpl struct {
	walletsUsecase wallets.WalletsUsecase
}

type WalletsController interface {
	InitWallets(ctx context.Context, req service.InitWalletRequest) (resp service.InitWalletResponse, err error)
	EnabledWallets(ctx context.Context, req service.EnableWalletRequest) (resp service.EnableWalletResponse, err error)
	ViewWalletsBalance(ctx context.Context, req service.ViewWalletBalanceRequest) (resp service.ViewWalletBalanceResponse, err error)
	DisabledWallets(ctx context.Context, req service.DisableWalletRequest) (resp service.DisableWalletResponse, err error)
}

func NewWalletsController(walletsUsecase wallets.WalletsUsecase) WalletsController {
	return &WalletsControllerImpl{walletsUsecase: walletsUsecase}
}

func (wc *WalletsControllerImpl) InitWallets(ctx context.Context, req service.InitWalletRequest) (resp service.InitWalletResponse, err error) {

	req.ID = uuid.NewV4().String()
	req.Token = uuid.NewV4().String()

	return wc.walletsUsecase.InitWallets(ctx, req)
}

func (wc *WalletsControllerImpl) EnabledWallets(ctx context.Context, req service.EnableWalletRequest) (resp service.EnableWalletResponse, err error) {
	return wc.walletsUsecase.EnabledWallets(ctx, req)
}

func (wc *WalletsControllerImpl) ViewWalletsBalance(ctx context.Context, req service.ViewWalletBalanceRequest) (resp service.ViewWalletBalanceResponse, err error) {
	return wc.walletsUsecase.ViewWalletsBalance(ctx, req)
}

func (wc *WalletsControllerImpl) DisabledWallets(ctx context.Context, req service.DisableWalletRequest) (resp service.DisableWalletResponse, err error) {
	return wc.walletsUsecase.DisabledWallets(ctx, req)
}

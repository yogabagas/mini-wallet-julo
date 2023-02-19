package controller

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/yogabagas/jatis-BE/domain/service"
	wallets "github.com/yogabagas/jatis-BE/service/wallets/usecase"
)

type WalletsControllerImpl struct {
	walletsUsecase wallets.WalletsUsecase
}

type WalletsController interface {
	CreateWallets(ctx context.Context, req service.InitWalletRequest) (resp service.InitWalletResponse, err error)
}

func NewWalletsController(walletsUsecase wallets.WalletsUsecase) WalletsController {
	return &WalletsControllerImpl{walletsUsecase: walletsUsecase}
}

func (cc *WalletsControllerImpl) CreateWallets(ctx context.Context, req service.InitWalletRequest) (resp service.InitWalletResponse, err error) {

	req.ID = uuid.NewV4().String()
	req.Token = uuid.NewV4().String()

	return cc.walletsUsecase.CreateWallets(ctx, req)
}

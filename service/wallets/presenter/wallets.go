package presenter

import (
	"context"
	"errors"

	"github.com/yogabagas/jatis-BE/domain/model"
	"github.com/yogabagas/jatis-BE/domain/service"
)

type WalletsPresenterImpl struct{}

type WalletsPresenter interface {
	CreateWallet(ctx context.Context, req model.Wallets) (resp service.InitWalletResponse, err error)
}

func NewWalletsPresenter() WalletsPresenter {
	return &WalletsPresenterImpl{}
}

func (wp *WalletsPresenterImpl) CreateWallet(ctx context.Context, req model.Wallets) (resp service.InitWalletResponse, err error) {
	if req.Token == "" {
		return resp, errors.New("token is nil")
	}
	return service.InitWalletResponse{Token: req.Token}, nil
}

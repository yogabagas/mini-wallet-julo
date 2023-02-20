package presenter

import (
	"context"
	"errors"
	"time"

	"github.com/yogabagas/mini-wallet-julo/domain/model"
	"github.com/yogabagas/mini-wallet-julo/domain/service"
	"github.com/yogabagas/mini-wallet-julo/shared/constant"
)

type WalletsPresenterImpl struct{}

type WalletsPresenter interface {
	InitWallets(ctx context.Context, req model.CreateWalletRequest) (resp service.InitWalletResponse, err error)
	EnabledWallets(ctx context.Context, req model.ReadWalletByTokenResponse) (resp service.EnableWalletResponse, err error)
	ViewWalletsBalance(ctx context.Context, req model.ReadWalletByTokenResponse) (resp service.ViewWalletBalanceResponse, err error)
	DisabledWallets(ctx context.Context, req model.ReadWalletByTokenResponse) (resp service.DisableWalletResponse, err error)
}

func NewWalletsPresenter() WalletsPresenter {
	return &WalletsPresenterImpl{}
}

func (wp *WalletsPresenterImpl) InitWallets(ctx context.Context, req model.CreateWalletRequest) (resp service.InitWalletResponse, err error) {
	if req.Token == "" {
		return resp, errors.New("token is nil")
	}
	return service.InitWalletResponse{Token: req.Token}, nil
}

func (wp *WalletsPresenterImpl) EnabledWallets(ctx context.Context, req model.ReadWalletByTokenResponse) (resp service.EnableWalletResponse, err error) {

	if req.ID == "" {
		return resp, errors.New("wallet not found")
	}

	now := time.Now()

	return service.EnableWalletResponse{
		Wallet: service.Wallet{
			ID:        req.ID,
			OwnedBy:   req.OwnedBy,
			Status:    constant.WalletStatus(req.Status).String(),
			EnabledAt: &now,
			Balance:   int(req.Balance),
		},
	}, nil

}

func (wp *WalletsPresenterImpl) ViewWalletsBalance(ctx context.Context, req model.ReadWalletByTokenResponse) (resp service.ViewWalletBalanceResponse, err error) {

	if req.ID == "" {
		return resp, errors.New("wallet not found")
	}

	if req.Status == constant.Disabled.Int() {
		return service.ViewWalletBalanceResponse{
			Error: "Wallet disabled",
		}, nil
	}

	now := time.Now()

	return service.ViewWalletBalanceResponse{
		Wallet: &service.Wallet{
			ID:        req.ID,
			OwnedBy:   req.OwnedBy,
			Status:    constant.WalletStatus(req.Status).String(),
			EnabledAt: &now,
			Balance:   int(req.Balance),
		},
	}, nil
}

func (wp *WalletsPresenterImpl) DisabledWallets(ctx context.Context, req model.ReadWalletByTokenResponse) (resp service.DisableWalletResponse, err error) {

	if req.ID == "" {
		return resp, errors.New("wallet not found")
	}

	now := time.Now()

	return service.DisableWalletResponse{
		Wallet: service.Wallet{
			ID:         req.ID,
			OwnedBy:    req.OwnedBy,
			Status:     constant.WalletStatus(req.Status).String(),
			DisabledAt: &now,
			Balance:    int(req.Balance),
		},
	}, nil
}

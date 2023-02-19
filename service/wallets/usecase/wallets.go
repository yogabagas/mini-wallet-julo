package usecase

import (
	"context"
	"time"

	"github.com/yogabagas/jatis-BE/domain/model"
	"github.com/yogabagas/jatis-BE/domain/repository"
	"github.com/yogabagas/jatis-BE/domain/service"
	"github.com/yogabagas/jatis-BE/service/wallets/presenter"
	"github.com/yogabagas/jatis-BE/shared/constant"
)

type WalletsUsecaseImpl struct {
	repoRegistry repository.RepositoryRegistry
	presenter    presenter.WalletsPresenter
}

type WalletsUsecase interface {
	CreateWallets(ctx context.Context, req service.InitWalletRequest) (resp service.InitWalletResponse, err error)
}

func NewWalletsUsecase(repoRegistry repository.RepositoryRegistry, presenter presenter.WalletsPresenter) WalletsUsecase {
	return &WalletsUsecaseImpl{
		repoRegistry: repoRegistry,
		presenter:    presenter}
}

func (wu *WalletsUsecaseImpl) CreateWallets(ctx context.Context, req service.InitWalletRequest) (resp service.InitWalletResponse, err error) {

	walletRepository := wu.repoRegistry.GetWalletsRepository()

	reqWallet := model.Wallets{
		ID:        req.ID,
		OwnedBy:   req.CustomerID,
		Token:     req.Token,
		Status:    constant.Disabled.Int(),
		Balance:   0,
		EnabledAt: time.Time{},
	}

	err = walletRepository.CreateWallet(ctx, reqWallet)
	if err != nil {
		return resp, err
	}

	return wu.presenter.CreateWallet(ctx, reqWallet)
}

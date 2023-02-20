package presenter

import (
	"context"
	"errors"
	"time"

	"github.com/yogabagas/mini-wallet-julo/domain/model"
	"github.com/yogabagas/mini-wallet-julo/domain/service"
	"github.com/yogabagas/mini-wallet-julo/shared/constant"
)

type WalletBalanceHistoriesPresenterImpl struct{}

type WalletBalanceHistoriesPresenter interface {
	Deposit(ctx context.Context, amount float64, refID string, req model.ReadWalletByTokenResponse) (resp service.WalletDepositResponse, err error)
	Withdrawal(ctx context.Context, amount float64, refID string, req model.ReadWalletByTokenResponse) (resp service.WalletWithdrawalsResponse, err error)
	Transactions(ctx context.Context, req model.ReadWalletBalanceHisoriesByWalletIDResponse) (resp service.WalletTransactionsResponse, err error)
}

func NewWalletBalanceHistoriesPresenter() WalletBalanceHistoriesPresenter {
	return &WalletBalanceHistoriesPresenterImpl{}
}

func (wp *WalletBalanceHistoriesPresenterImpl) Deposit(ctx context.Context, amount float64, refID string, req model.ReadWalletByTokenResponse) (resp service.WalletDepositResponse, err error) {

	if req.ID == "" {
		return resp, errors.New("wallet not found")
	}

	return service.WalletDepositResponse{
		Deposit: service.Deposit{
			ID:          req.ID,
			DepositBy:   req.OwnedBy,
			Status:      "success",
			DepositAt:   time.Now(),
			Amount:      int(amount),
			ReferenceID: refID,
		},
	}, nil
}

func (wp *WalletBalanceHistoriesPresenterImpl) Withdrawal(ctx context.Context, amount float64, refID string, req model.ReadWalletByTokenResponse) (resp service.WalletWithdrawalsResponse, err error) {

	if req.ID == "" {
		return resp, errors.New("wallet not found")
	}

	return service.WalletWithdrawalsResponse{
		Withdrawal: service.Withdrawal{
			ID:           req.ID,
			DepositBy:    req.OwnedBy,
			Status:       "success",
			WithdrawalAt: time.Now(),
			Amount:       int(amount),
			ReferenceID:  refID,
		},
	}, nil
}

func (wp *WalletBalanceHistoriesPresenterImpl) Transactions(ctx context.Context, req model.ReadWalletBalanceHisoriesByWalletIDResponse) (resp service.WalletTransactionsResponse, err error) {

	if len(req.WalletBalanceHistories) > 0 {

		for _, v := range req.WalletBalanceHistories {
			history := service.WalletBalanceHistory{
				WalletID:    v.WalletID,
				ReferenceID: v.ReferenceID,
				Amount:      int(v.Amount),
				Type:        constant.WalletType(v.Type).String(),
				Description: v.Description,
				CreatedAt:   v.CreatedAt,
				CreatedBy:   v.CreatedBy,
				UpdatedAt:   v.UpdatedAt,
				UpdatedBy:   v.UpdatedBy,
			}
			resp.Histories = append(resp.Histories, history)
		}
	}

	return resp, nil
}

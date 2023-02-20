package repository

import (
	"context"

	"github.com/yogabagas/mini-wallet-julo/domain/model"
)

const (
	insertWalletBalanceHistories = `INSERT INTO wallet_balance_histories (wallet_id, reference_id, amount, type, description, created_by, updated_by) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	selectSumAmountByWalletID = `SELECT SUM(amount) AS total_amount, wallet_id FROM wallet_balance_histories WHERE wallet_id = ? GROUP BY wallet_id`

	selectHistoryByWalletID = `SELECT wallet_id, reference_id, amount, type, description, created_at, created_by, updated_at, updated_by FROM wallet_balance_histories
	WHERE wallet_id = ? ORDER BY created_at DESC`
)

type WalletBalanceHistoriesRepositoryImpl struct {
	db DBExecutor
}

type WalletBalanceHistoriesRepository interface {
	InsertWalletBalanceHistories(ctx context.Context, req model.InsertWalletBalanceHistoriesRequest) error
	ReadSumAmountByWalletID(ctx context.Context, req model.ReadSumAmountByWalletIDRequest) (resp model.ReadSumAmountByWalletIDResponse, err error)
	ReadWalletBalanceHisoriesByWalletID(ctx context.Context, req model.ReadWalletBalanceHisoriesByWalletIDRequest) (resp model.ReadWalletBalanceHisoriesByWalletIDResponse, err error)
}

func NewWalletBalanceHistoriesRepository(db DBExecutor) WalletBalanceHistoriesRepository {
	return &WalletBalanceHistoriesRepositoryImpl{db: db}
}

func (wr *WalletBalanceHistoriesRepositoryImpl) InsertWalletBalanceHistories(ctx context.Context, req model.InsertWalletBalanceHistoriesRequest) error {

	_, err := wr.db.ExecContext(ctx, insertWalletBalanceHistories, req.WalletID, req.ReferenceID, req.Amount, req.Type, req.Description, req.CreatedBy, req.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (wr *WalletBalanceHistoriesRepositoryImpl) ReadSumAmountByWalletID(ctx context.Context, req model.ReadSumAmountByWalletIDRequest) (resp model.ReadSumAmountByWalletIDResponse, err error) {

	err = wr.db.QueryRowContext(ctx, selectSumAmountByWalletID, req.WalletID).Scan(&resp.TotalAmount, &resp.WalletID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (wr *WalletBalanceHistoriesRepositoryImpl) ReadWalletBalanceHisoriesByWalletID(ctx context.Context, req model.ReadWalletBalanceHisoriesByWalletIDRequest) (resp model.ReadWalletBalanceHisoriesByWalletIDResponse, err error) {

	rows, err := wr.db.QueryContext(ctx, selectHistoryByWalletID, req.WalletID)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		wbHistory := model.WalletBalanceHistories{}

		err = rows.Scan(&wbHistory.WalletID, &wbHistory.ReferenceID, &wbHistory.Amount, &wbHistory.Type,
			&wbHistory.Description, &wbHistory.CreatedAt, &wbHistory.CreatedBy, &wbHistory.UpdatedAt, &wbHistory.UpdatedBy)
		if err != nil {
			return resp, err
		}

		resp.WalletBalanceHistories = append(resp.WalletBalanceHistories, wbHistory)
	}

	return resp, nil
}

package repository

import (
	"context"
	"database/sql"

	"github.com/yogabagas/mini-wallet-julo/domain/model"
)

const (
	insertWallet = `INSERT INTO wallets (id, owned_by, token, status, balance, enabled_at, created_by, updated_by)
	VALUES(?,?,?,?,?,?,?,?,?,?)`

	updateWalletByToken = `UPDATE wallets SET status = ?, enabled_at = ?, updated_by = ? WHERE token = ?`

	updateWalletDisableByToken = `UPDATE wallets SET status = ?, updated_by = ? WHERE token = ?`

	selectWalletByToken = `SELECT id, owned_by, token, status, balance, enabled_at, created_at, created_by, updated_at, updated_by FROM wallets
	WHERE token = ?`

	updateAmountWalletByID = `UPDATE wallets w 
	SET balance = (SELECT SUM(amount) FROM wallet_balance_histories WHERE wallet_id = w.id) WHERE w.id = ?`
)

type WalletsRepositoryImpl struct {
	db DBExecutor
}

type WalletsRepository interface {
	CreateWallet(ctx context.Context, req model.CreateWalletRequest) error
	UpdateWalletByToken(ctx context.Context, req model.UpdateWalletByTokenRequest) error
	ReadWalletByToken(ctx context.Context, req model.ReadWalletByTokenRequest) (resp model.ReadWalletByTokenResponse, err error)
	UpdateAmountWalletByID(ctx context.Context, req model.UpdateAmountWalletByIDRequest) error
}

func NewWalletsRepository(db DBExecutor) WalletsRepository {
	return &WalletsRepositoryImpl{db: db}
}

func (wr *WalletsRepositoryImpl) CreateWallet(ctx context.Context, req model.CreateWalletRequest) error {

	_, err := wr.db.ExecContext(ctx, insertWallet, req.ID, req.OwnedBy, req.Token, req.Status, req.Balance, req.CreatedBy, req.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (wr *WalletsRepositoryImpl) UpdateWalletByToken(ctx context.Context, req model.UpdateWalletByTokenRequest) error {

	if req.EnabledAt.IsZero() {
		_, err := wr.db.ExecContext(ctx, updateWalletDisableByToken, req.Status, req.UpdatedBy, req.Token)
		if err != nil {
			return err
		}
	} else {
		_, err := wr.db.ExecContext(ctx, updateWalletByToken, req.Status, req.EnabledAt, req.UpdatedBy, req.Token)
		if err != nil {
			return err
		}
	}

	return nil
}

func (wr *WalletsRepositoryImpl) ReadWalletByToken(ctx context.Context, req model.ReadWalletByTokenRequest) (resp model.ReadWalletByTokenResponse, err error) {

	nullTime := sql.NullTime{}

	err = wr.db.QueryRowContext(ctx, selectWalletByToken, req.Token).
		Scan(&resp.ID, &resp.OwnedBy, &resp.Token, &resp.Status, &resp.Balance, &nullTime,
			&resp.CreatedAt, &resp.CreatedBy, &resp.UpdatedAt, &resp.UpdatedBy)
	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	resp.EnabledAt = nullTime.Time

	return resp, nil
}

func (wr *WalletsRepositoryImpl) UpdateAmountWalletByID(ctx context.Context, req model.UpdateAmountWalletByIDRequest) error {

	_, err := wr.db.ExecContext(ctx, updateAmountWalletByID, req.ID)
	if err != nil {
		return err
	}
	return nil
}

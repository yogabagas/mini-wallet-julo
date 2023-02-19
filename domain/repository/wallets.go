package repository

import (
	"context"
	"time"

	"github.com/yogabagas/jatis-BE/domain/model"
)

const (
	insertWallet = `INSERT INTO wallets (id, owned_by, token, status, balance, enabled_at, created_at, created_by, updated_at, updated_by)
	VALUES(?,?,?,?,?,?,?,?,?)`
)

type WalletsRepositoryImpl struct {
	db DBExecutor
}

type WalletsRepository interface {
	CreateWallet(ctx context.Context, req model.Wallets) error
}

func NewWalletsRepository(db DBExecutor) WalletsRepository {
	return &WalletsRepositoryImpl{db: db}
}

func (wr *WalletsRepositoryImpl) CreateWallet(ctx context.Context, req model.Wallets) error {

	now := time.Now()
	_, err := wr.db.ExecContext(ctx, insertWallet, req.ID, req.OwnedBy, req.Status, req.Balance, req.EnabledAt, now, req.CreatedBy, now, req.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}

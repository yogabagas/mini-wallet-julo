package repository

import (
	"context"
	"database/sql"
)

type InTransaction func(RepositoryRegistry) (interface{}, error)

type RepositoryRegistryImpl struct {
	db         *sql.DB
	dbExecutor DBExecutor
}

type RepositoryRegistry interface {
	GetWalletsRepository() WalletsRepository
	GetWalletBalanceHistoriesRepository() WalletBalanceHistoriesRepository

	DoInTransaction(ctx context.Context, txFunc InTransaction) (out interface{}, err error)
}

func NewRepositoryRegistry(db *sql.DB) RepositoryRegistry {
	return &RepositoryRegistryImpl{db: db}
}

func (r RepositoryRegistryImpl) GetWalletsRepository() WalletsRepository {
	if r.dbExecutor != nil {
		return NewWalletsRepository(r.dbExecutor)
	}
	return NewWalletsRepository(r.db)
}

func (r RepositoryRegistryImpl) GetWalletBalanceHistoriesRepository() WalletBalanceHistoriesRepository {
	if r.dbExecutor != nil {
		return NewWalletBalanceHistoriesRepository(r.dbExecutor)
	}
	return NewWalletBalanceHistoriesRepository(r.db)
}

func (r RepositoryRegistryImpl) DoInTransaction(ctx context.Context, txFunc InTransaction) (out interface{}, err error) {
	var tx *sql.Tx

	registry := r

	if r.dbExecutor == nil {
		tx, err = r.db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p)
			} else if err != nil {
				rErr := tx.Rollback()
				if rErr != nil {
					err = rErr
				}
			} else {
				err = tx.Commit()
			}
		}()
		registry = RepositoryRegistryImpl{
			db:         r.db,
			dbExecutor: tx,
		}
	}
	out, err = txFunc(registry)
	return
}

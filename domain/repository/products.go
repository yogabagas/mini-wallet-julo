package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/yogabagas/jatis-BE/domain/model"
)

const (
	insertProducts = `INSERT INTO products (product_name, unit_price, in_stock, is_deleted, created_at, created_by, updated_at, updated_by) VALUES %s`
)

type ProductsRepositoryImpl struct {
	db *sql.DB
}

type ProductsRepository interface {
	InsertProducts(ctx context.Context, req []model.Product) error
}

func NewProductsRepository(db *sql.DB) ProductsRepository {
	return &ProductsRepositoryImpl{db: db}
}

func (pr *ProductsRepositoryImpl) InsertProducts(ctx context.Context, req []model.Product) error {
	tx, err := pr.db.Begin()
	if err != nil {
		return err
	}

	valueStrings := []string{}
	valuesArgs := []interface{}{}

	now := time.Now()

	if len(req) > 0 {
		for _, v := range req {
			valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?)")
			valuesArgs = append(valuesArgs, v.ProductName)
			valuesArgs = append(valuesArgs, v.UnitPrice)
			valuesArgs = append(valuesArgs, v.InStock)
			valuesArgs = append(valuesArgs, false)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.CreatedBy)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.UpdatedBy)
		}
	}

	q := fmt.Sprintf(insertProducts, strings.Join(valueStrings, ","))
	_, err = tx.Exec(q, valuesArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

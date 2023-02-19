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
	insertShippingMethods = `INSERT INTO shipping_methods (shipping_method, is_deleted, created_at, created_by, updated_at, updated_by) VALUES %s`
)

type ShippingMethodsRepositoryImpl struct {
	db *sql.DB
}

type ShippingMethodsRepository interface {
	InsertShippingMethods(ctx context.Context, req []model.ShippingMethod) error
}

func NewShippingMethodsRepository(db *sql.DB) ShippingMethodsRepository {
	return &ShippingMethodsRepositoryImpl{db: db}
}

func (er *ShippingMethodsRepositoryImpl) InsertShippingMethods(ctx context.Context, req []model.ShippingMethod) error {
	tx, err := er.db.Begin()
	if err != nil {
		return err
	}

	valueStrings := []string{}
	valuesArgs := []interface{}{}

	now := time.Now()

	if len(req) > 0 {
		for _, v := range req {
			valueStrings = append(valueStrings, "(?,?,?,?,?,?)")
			valuesArgs = append(valuesArgs, v.ShippingMethodName)
			valuesArgs = append(valuesArgs, false)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.CreatedBy)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.UpdatedBy)
		}
	}

	q := fmt.Sprintf(insertShippingMethods, strings.Join(valueStrings, ","))
	_, err = tx.Exec(q, valuesArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

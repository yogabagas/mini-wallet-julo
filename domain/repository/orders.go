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
	insertOrders = `INSERT INTO orders (customer_id, employee_id, order_date, purchase_order_number, ship_date, shipping_method_id, freight_charge, 
		taxes, payment_received, comment, created_at, created_by, updated_at, updated_by) VALUES %s`
)

type OrdersRepositoryImpl struct {
	db *sql.DB
}

type OrdersRepository interface {
	InsertOrders(ctx context.Context, req []model.Order) error
}

func NewOrdersRepository(db *sql.DB) OrdersRepository {
	return &OrdersRepositoryImpl{db: db}
}

func (er *OrdersRepositoryImpl) InsertOrders(ctx context.Context, req []model.Order) error {
	tx, err := er.db.Begin()
	if err != nil {
		return err
	}

	valueStrings := []string{}
	valuesArgs := []interface{}{}

	now := time.Now()

	if len(req) > 0 {
		for _, v := range req {
			valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
			valuesArgs = append(valuesArgs, v.CustomerId)
			valuesArgs = append(valuesArgs, v.EmployeeId)
			valuesArgs = append(valuesArgs, v.OrderDate)
			valuesArgs = append(valuesArgs, v.PurchaseOrderNumber)
			valuesArgs = append(valuesArgs, v.ShipDate)
			valuesArgs = append(valuesArgs, v.ShippingMethodId)
			valuesArgs = append(valuesArgs, v.FreightCharge)
			valuesArgs = append(valuesArgs, v.Taxes)
			valuesArgs = append(valuesArgs, v.PaymentReceived)
			valuesArgs = append(valuesArgs, v.Comment)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.CreatedBy)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.UpdatedBy)
		}
	}

	q := fmt.Sprintf(insertOrders, strings.Join(valueStrings, ","))
	_, err = tx.Exec(q, valuesArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

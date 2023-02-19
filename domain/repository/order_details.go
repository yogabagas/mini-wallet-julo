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
	insertOrderDetails = `INSERT INTO order_details (order_id, product_id, qty, unit_price, discount, created_at, created_by, updated_at, updated_by) VALUES %s`

	selectOrderDetails = `SELECT CONCAT(c.first_name, " ", c.last_name) AS customer_name, CONCAT(e.first_name, " ", e.last_name) AS employee_name, sm.shipping_method,  od.qty, od.unit_price, od.discount, (od.qty * od.unit_price) - (od.qty * od.unit_price *
		(od.discount/100)) AS subtotal_price, ((od.qty * od.unit_price) - (od.qty * od.unit_price *
		(od.discount/100))+p.unit_price) AS total_payment FROM orders o 
		LEFT JOIN customers c ON o.customer_id = c.customer_id LEFT JOIN employees e ON o.employee_id = e.employee_id 
		LEFT JOIN shipping_methods sm ON o.shipping_method_id = sm.shipping_method_id LEFT JOIN order_details od ON o.order_id = od.order_id  LEFT JOIN products p ON od.product_id = p.product_id;`
)

type OrderDetailsRepositoryImpl struct {
	db *sql.DB
}

type OrderDetaisRepository interface {
	CreateOrderDetails(ctx context.Context, req []model.OrderDetail) error
	ReadOrderDetails(context.Context) ([]model.ReadOrderDetailsResp, error)
}

func NewOrderDetailsRepository(db *sql.DB) OrderDetaisRepository {
	return &OrderDetailsRepositoryImpl{db: db}
}

func (or *OrderDetailsRepositoryImpl) CreateOrderDetails(ctx context.Context, req []model.OrderDetail) error {
	tx, err := or.db.Begin()
	if err != nil {
		return err
	}

	valueStrings := []string{}
	valuesArgs := []interface{}{}

	now := time.Now()

	if len(req) > 0 {
		for _, v := range req {
			valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?,?)")
			valuesArgs = append(valuesArgs, v.OrderId)
			valuesArgs = append(valuesArgs, v.ProductId)
			valuesArgs = append(valuesArgs, v.Qty)
			valuesArgs = append(valuesArgs, v.UnitPrice)
			valuesArgs = append(valuesArgs, v.Discount)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.CreatedBy)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.UpdatedBy)
		}
	}

	q := fmt.Sprintf(insertOrderDetails, strings.Join(valueStrings, ","))
	_, err = tx.Exec(q, valuesArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (or *OrderDetailsRepositoryImpl) ReadOrderDetails(context.Context) (resp []model.ReadOrderDetailsResp, err error) {

	rows, err := or.db.Query(selectOrderDetails)
	if err != nil {
		return
	}

	for rows.Next() {
		res := model.ReadOrderDetailsResp{}

		err = rows.Scan(&res.CustomerName, &res.EmployeeName, &res.ShippingMethod, &res.Qty, &res.UnitPrice, &res.Discount, &res.SubTotalPrice, &res.TotalPayment)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		resp = append(resp, res)
	}

	return resp, nil
}

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
	insertCustomers = `INSERT INTO customers (company_name, first_name, last_name, billing_address, province, zip_code, email, company_website,
    phone_number, fax_number, ship_address, ship_city, ship_province, ship_zip_code, ship_phone_number, is_deleted, created_at, created_by, updated_at, updated_by) VALUES %s`
)

type CustomersRepositoryImpl struct {
	db *sql.DB
}

type CustomersRepository interface {
	InsertCustomers(ctx context.Context, req []model.Customer) error
}

func NewCustomersRepository(db *sql.DB) CustomersRepository {
	return &CustomersRepositoryImpl{db: db}
}

func (cr *CustomersRepositoryImpl) InsertCustomers(ctx context.Context, req []model.Customer) error {

	tx, err := cr.db.Begin()
	if err != nil {
		return err
	}

	valueStrings := []string{}
	valuesArgs := []interface{}{}

	now := time.Now()

	if len(req) > 0 {
		for _, v := range req {
			valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
			valuesArgs = append(valuesArgs, v.CompanyName)
			valuesArgs = append(valuesArgs, v.FirstName)
			valuesArgs = append(valuesArgs, v.LastName)
			valuesArgs = append(valuesArgs, v.BillingAddress)
			valuesArgs = append(valuesArgs, v.Province)
			valuesArgs = append(valuesArgs, v.ZipCode)
			valuesArgs = append(valuesArgs, v.Email)
			valuesArgs = append(valuesArgs, v.CompanyWebsite)
			valuesArgs = append(valuesArgs, v.PhoneNumber)
			valuesArgs = append(valuesArgs, v.FaxNumber)
			valuesArgs = append(valuesArgs, v.ShipAddress)
			valuesArgs = append(valuesArgs, v.ShipCity)
			valuesArgs = append(valuesArgs, v.ShipProvince)
			valuesArgs = append(valuesArgs, v.ShipZipCode)
			valuesArgs = append(valuesArgs, v.ShipPhoneNumber)
			valuesArgs = append(valuesArgs, false)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.CreatedBy)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.UpdatedBy)
		}
	}

	q := fmt.Sprintf(insertCustomers, strings.Join(valueStrings, ","))
	_, err = tx.Exec(q, valuesArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

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
	insertEmployees = `INSERT INTO employees (first_name, last_name, title, work_phone, is_deleted, created_at, created_by, updated_at, updated_by) VALUES %s`
)

type EmployeesRepositoryImpl struct {
	db *sql.DB
}

type EmployeesRepository interface {
	InsertEmployees(ctx context.Context, req []model.Employee) error
}

func NewEmployeesRepository(db *sql.DB) EmployeesRepository {
	return &EmployeesRepositoryImpl{db: db}
}

func (er *EmployeesRepositoryImpl) InsertEmployees(ctx context.Context, req []model.Employee) error {
	tx, err := er.db.Begin()
	if err != nil {
		return err
	}

	valueStrings := []string{}
	valuesArgs := []interface{}{}

	now := time.Now()

	if len(req) > 0 {
		for _, v := range req {
			valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?,?)")
			valuesArgs = append(valuesArgs, v.FirstName)
			valuesArgs = append(valuesArgs, v.LastName)
			valuesArgs = append(valuesArgs, v.Title)
			valuesArgs = append(valuesArgs, v.WorkPhone)
			valuesArgs = append(valuesArgs, false)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.CreatedBy)
			valuesArgs = append(valuesArgs, now)
			valuesArgs = append(valuesArgs, v.UpdatedBy)
		}
	}

	q := fmt.Sprintf(insertEmployees, strings.Join(valueStrings, ","))
	_, err = tx.Exec(q, valuesArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

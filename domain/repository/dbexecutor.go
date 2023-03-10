package repository

import (
	"context"
	"database/sql"
)

type DBExecutor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	// SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	// NamedExecContext(ctx context.Context, query string, args interface{}) (sql.Result, error)
	// Rebind(query string) string
}

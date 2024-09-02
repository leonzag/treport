package database

import (
	"context"
	"database/sql"
)

type DB interface {
	Close() error

	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	PingContext(ctx context.Context) error

	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

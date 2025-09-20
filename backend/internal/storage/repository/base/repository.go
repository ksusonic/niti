package base

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type contextKey string

const txKey contextKey = "tx"

type Repository struct {
	*pgxpool.Pool
}

func NewBaseRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{Pool: pool}
}

// Exec executes a query using transaction if available in context
func (r *Repository) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	if tx, ok := ctx.Value(txKey).(pgx.Tx); ok {
		return tx.Exec(ctx, sql, arguments...)
	}
	return r.Pool.Exec(ctx, sql, arguments...)
}

// Query executes a query using transaction if available in context
func (r *Repository) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	if tx, ok := ctx.Value(txKey).(pgx.Tx); ok {
		return tx.Query(ctx, sql, args...)
	}
	return r.Pool.Query(ctx, sql, args...)
}

// QueryRow executes a query that returns a single row using transaction if available in context
func (r *Repository) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	if tx, ok := ctx.Value(txKey).(pgx.Tx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return r.Pool.QueryRow(ctx, sql, args...)
}

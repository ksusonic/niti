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
	pool *pgxpool.Pool
}

func NewBaseRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

// Exec executes a query using transaction if available in context
func (r *Repository) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	if tx, ok := ctx.Value(txKey).(pgx.Tx); ok {
		return tx.Exec(ctx, sql, arguments...)
	}
	return r.pool.Exec(ctx, sql, arguments...)
}

// Query executes a query using transaction if available in context
func (r *Repository) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	if tx, ok := ctx.Value(txKey).(pgx.Tx); ok {
		return tx.Query(ctx, sql, args...)
	}
	return r.pool.Query(ctx, sql, args...)
}

// QueryRow executes a query that returns a single row using transaction if available in context
func (r *Repository) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if tx, ok := ctx.Value(txKey).(pgx.Tx); ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return r.pool.QueryRow(ctx, sql, args...)
}

// Acquire acquires a connection from the pool
func (r *Repository) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	return r.pool.Acquire(ctx)
}

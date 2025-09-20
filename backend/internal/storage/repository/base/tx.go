package base

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) WithTx(ctx context.Context, fn func(ctx context.Context) error) (err error) {
	conn, err := r.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err2 := tx.Rollback(ctx)
			if err2 != nil {
				err = fmt.Errorf("rollback: %w, original error: %w", err2, err)
			}
		} else {
			err2 := tx.Commit(ctx)
			if err2 != nil {
				err = fmt.Errorf("commit: %w", err2)
			}
		}
	}()

	err = fn(context.WithValue(ctx, txKey, tx))

	return
}

func (r *Repository) WithRollback(ctx context.Context, fn func(ctx context.Context)) (err error) {
	conn, err := r.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	fn(context.WithValue(ctx, txKey, tx))

	return
}

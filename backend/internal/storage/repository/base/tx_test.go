//go:build integration

package base_test

import (
	"context"
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/ksusonic/niti/backend/internal/storage/repository/base"
	"github.com/ksusonic/niti/backend/pkg/config"
	"github.com/stretchr/testify/require"
)

func setupTestPool(t *testing.T) *pgxpool.Pool {
	require.NoError(t, godotenv.Load("../../../../.env"))

	var cfg config.PostgresConfig
	require.NoError(t, env.Parse(&cfg))

	// Parse the DSN and configure to disable prepared statement caching
	config, err := pgxpool.ParseConfig(cfg.DSN)
	require.NoError(t, err)

	// Disable prepared statement caching to avoid conflicts in tests
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec

	pool, err := pgxpool.NewWithConfig(t.Context(), config)
	require.NoError(t, err)

	t.Cleanup(func() {
		pool.Close()
	})

	return pool
}

func createTestTable(t *testing.T, pool *pgxpool.Pool) {
	// Create a test table for our tests
	_, err := pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS test_tx_table (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			value INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)
	`)
	require.NoError(t, err)

	// Clean up any existing data
	_, err = pool.Exec(context.Background(), `DELETE FROM test_tx_table`)
	require.NoError(t, err)

	t.Cleanup(func() {
		pool.Exec(context.Background(), `DROP TABLE IF EXISTS test_tx_table`)
	})
}

func TestRepository_WithTx_Success(t *testing.T) {
	pool := setupTestPool(t)
	createTestTable(t, pool)

	repo := base.NewBaseRepository(pool)
	ctx := context.Background()

	var insertedID int
	err := repo.WithTx(ctx, func(ctx context.Context) error {
		// Insert a record
		row := repo.QueryRow(ctx,
			`INSERT INTO test_tx_table (name, value) VALUES ($1, $2) RETURNING id`,
			"test_name", 42)

		if err := row.Scan(&insertedID); err != nil {
			return err
		}

		// Update the record
		_, err := repo.Exec(ctx,
			`UPDATE test_tx_table SET value = $1 WHERE id = $2`,
			100, insertedID)

		return err
	})

	require.NoError(t, err)
	require.NotZero(t, insertedID)

	// Verify the transaction was committed
	var name string
	var value int
	row := pool.QueryRow(ctx, `SELECT name, value FROM test_tx_table WHERE id = $1`, insertedID)
	err = row.Scan(&name, &value)
	require.NoError(t, err)
	require.Equal(t, "test_name", name)
	require.Equal(t, 100, value)
}

func TestRepository_WithTx_Rollback_OnError(t *testing.T) {
	pool := setupTestPool(t)
	createTestTable(t, pool)

	repo := base.NewBaseRepository(pool)
	ctx := context.Background()

	// First, insert a record to verify rollback
	_, err := pool.Exec(ctx, `INSERT INTO test_tx_table (name, value) VALUES ($1, $2)`, "existing", 1)
	require.NoError(t, err)

	// This transaction should fail and rollback
	err = repo.WithTx(ctx, func(ctx context.Context) error {
		// Insert a new record
		_, err := repo.Exec(ctx,
			`INSERT INTO test_tx_table (name, value) VALUES ($1, $2)`,
			"should_rollback", 999)
		if err != nil {
			return err
		}

		// Update the existing record
		_, err = repo.Exec(ctx,
			`UPDATE test_tx_table SET value = $1 WHERE name = $2`,
			555, "existing")
		if err != nil {
			return err
		}

		// Force an error to trigger rollback
		return &customError{message: "intentional error"}
	})

	// Verify the error was returned
	require.Error(t, err)
	require.Contains(t, err.Error(), "intentional error")

	// Verify that no changes were committed (rollback occurred)
	var count int
	row := pool.QueryRow(ctx, `SELECT COUNT(*) FROM test_tx_table WHERE name = 'should_rollback'`)
	err = row.Scan(&count)
	require.NoError(t, err)
	require.Equal(t, 0, count, "Record should not exist due to rollback")

	// Verify the existing record wasn't modified
	var value int
	row = pool.QueryRow(ctx, `SELECT value FROM test_tx_table WHERE name = 'existing'`)
	err = row.Scan(&value)
	require.NoError(t, err)
	require.Equal(t, 1, value, "Existing record should not be modified due to rollback")
}

func TestRepository_WithRollback_AlwaysRollsBack(t *testing.T) {
	pool := setupTestPool(t)
	createTestTable(t, pool)

	repo := base.NewBaseRepository(pool)
	ctx := context.Background()

	var insertedID int
	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// Insert a record
		row := repo.QueryRow(ctx,
			`INSERT INTO test_tx_table (name, value) VALUES ($1, $2) RETURNING id`,
			"rollback_test", 777)

		// This should succeed within the transaction
		scanErr := row.Scan(&insertedID)
		require.NoError(t, scanErr)
		require.NotZero(t, insertedID)

		// Update the record
		_, execErr := repo.Exec(ctx,
			`UPDATE test_tx_table SET value = $1 WHERE id = $2`,
			888, insertedID)
		require.NoError(t, execErr)

		// Verify the record exists within the transaction
		var count int
		countRow := repo.QueryRow(ctx, `SELECT COUNT(*) FROM test_tx_table WHERE id = $1`, insertedID)
		countErr := countRow.Scan(&count)
		require.NoError(t, countErr)
		require.Equal(t, 1, count)
	})

	require.NoError(t, err)

	// Verify the transaction was rolled back (record should not exist)
	var count int
	row := pool.QueryRow(ctx, `SELECT COUNT(*) FROM test_tx_table WHERE name = 'rollback_test'`)
	err = row.Scan(&count)
	require.NoError(t, err)
	require.Equal(t, 0, count, "Record should not exist due to rollback")
}

func TestRepository_WithRollback_WithError(t *testing.T) {
	pool := setupTestPool(t)
	createTestTable(t, pool)

	repo := base.NewBaseRepository(pool)
	ctx := context.Background()

	// Even if operations fail within WithRollback, it should handle gracefully
	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// Try to insert with a constraint violation or other error
		_, err := repo.Exec(ctx,
			`INSERT INTO test_tx_table (name, value) VALUES ($1, $2)`,
			"error_test", 123)
		// We don't require.NoError here because WithRollback doesn't return errors from the callback
		_ = err

		// Try to query non-existent data
		var nonExistent string
		row := repo.QueryRow(ctx, `SELECT name FROM test_tx_table WHERE id = -999`)
		err = row.Scan(&nonExistent)
		// Again, we don't handle the error since WithRollback doesn't propagate them
		_ = err
	})

	// WithRollback should always succeed unless there's an issue with the transaction itself
	require.NoError(t, err)

	// Verify no records were inserted
	var count int
	row := pool.QueryRow(ctx, `SELECT COUNT(*) FROM test_tx_table`)
	err = row.Scan(&count)
	require.NoError(t, err)
	require.Equal(t, 0, count)
}

func TestRepository_TransactionIsolation(t *testing.T) {
	pool := setupTestPool(t)
	createTestTable(t, pool)

	repo := base.NewBaseRepository(pool)
	ctx := context.Background()

	// Insert initial data
	_, err := pool.Exec(ctx, `INSERT INTO test_tx_table (name, value) VALUES ('isolation_test', 100)`)
	require.NoError(t, err)

	err = repo.WithTx(ctx, func(txCtx context.Context) error {
		// Update within transaction
		_, err := repo.Exec(txCtx,
			`UPDATE test_tx_table SET value = 200 WHERE name = 'isolation_test'`)
		if err != nil {
			return err
		}

		// Verify the change is visible within the transaction
		var valueInTx int
		row := repo.QueryRow(txCtx, `SELECT value FROM test_tx_table WHERE name = 'isolation_test'`)
		if err := row.Scan(&valueInTx); err != nil {
			return err
		}
		require.Equal(t, 200, valueInTx)

		// Verify the change is NOT visible outside the transaction yet
		var valueOutsideTx int
		outsideRow := pool.QueryRow(ctx, `SELECT value FROM test_tx_table WHERE name = 'isolation_test'`)
		if err := outsideRow.Scan(&valueOutsideTx); err != nil {
			return err
		}
		require.Equal(t, 100, valueOutsideTx, "Change should not be visible outside transaction")

		return nil
	})

	require.NoError(t, err)

	// After transaction commits, change should be visible
	var finalValue int
	row := pool.QueryRow(ctx, `SELECT value FROM test_tx_table WHERE name = 'isolation_test'`)
	err = row.Scan(&finalValue)
	require.NoError(t, err)
	require.Equal(t, 200, finalValue, "Change should be visible after commit")
}

func TestRepository_NestedTransactions(t *testing.T) {
	pool := setupTestPool(t)
	createTestTable(t, pool)

	repo := base.NewBaseRepository(pool)
	ctx := context.Background()

	err := repo.WithTx(ctx, func(outerCtx context.Context) error {
		// Insert in outer transaction
		_, err := repo.Exec(outerCtx,
			`INSERT INTO test_tx_table (name, value) VALUES ('outer', 1)`)
		if err != nil {
			return err
		}

		// Nested WithRollback creates its own transaction, separate from outer
		// This is the current behavior - it doesn't share the outer transaction
		err = repo.WithRollback(outerCtx, func(innerCtx context.Context) {
			// This runs in a separate transaction that will be rolled back
			_, err := repo.Exec(innerCtx,
				`INSERT INTO test_tx_table (name, value) VALUES ('inner', 2)`)
			require.NoError(t, err)
		})
		if err != nil {
			return err
		}

		// After WithRollback completes, only the outer transaction record should exist
		var count int
		row := repo.QueryRow(outerCtx, `SELECT COUNT(*) FROM test_tx_table`)
		err = row.Scan(&count)
		require.NoError(t, err)
		require.Equal(t, 1, count, "Only outer transaction record should exist")

		return nil
	})

	require.NoError(t, err)

	// Only the outer transaction record should exist since inner was rolled back
	var count int
	row := pool.QueryRow(ctx, `SELECT COUNT(*) FROM test_tx_table`)
	err = row.Scan(&count)
	require.NoError(t, err)
	require.Equal(t, 1, count, "Only outer record should exist, inner was rolled back")
}

// Custom error type for testing error handling
type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}

# Database Migrations

This directory contains PostgreSQL migrations for the NITI backend application using [golang-migrate](https://github.com/golang-migrate/migrate).

## Structure

- `migrations/` - Contains all migration files
- `internal/migrations/` - Go package for programmatic migration management
- `cmd/migrator/` - CLI tool for running migrations

## Migration Files

Migration files follow the naming convention:
```
{version}_{description}.up.sql
{version}_{description}.down.sql
```

Example:
```
001_create_refresh_tokens_table.up.sql
001_create_refresh_tokens_table.down.sql
```

## Configuration

Migrations are configured through environment variables:

- `POSTGRES_DSN` - Database connection string (required)
- `POSTGRES_MIGRATIONS_PATH` - Path to migrations directory (default: "migrations")

## Usage

### Using Justfile Commands

Run all pending migrations:
```bash
just migrate-up
```

Check current migration version:
```bash
just migrate-version
```

Create a new migration:
```bash
just migrate-create add_user_table
```

### Using the CLI Tool Directly

```bash
# Run migrations
cd backend
go run ./cmd/migrator/main.go -command up

# Check version
go run ./cmd/migrator/main.go -command version
```

### Programmatic Usage

```go
import (
    "github.com/ksusonic/niti/backend/internal/migrations"
    "github.com/ksusonic/niti/backend/pgk/config"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    // Run all pending migrations
    err = migrations.MigrateUp(cfg.Postgres)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Creating New Migrations

1. Install golang-migrate CLI:
   ```bash
   # macOS
   brew install golang-migrate
   
   # Linux
   curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
   sudo mv migrate /usr/local/bin/
   ```

2. Create migration files:
   ```bash
   just migrate-create create_users_table
   ```

3. Edit the generated `.up.sql` and `.down.sql` files with your SQL

## Best Practices

1. **Make migrations reversible** - Always provide a proper down migration
2. **Use transactions** - Wrap multiple statements in BEGIN/COMMIT
3. **Be idempotent** - Use IF EXISTS/IF NOT EXISTS where appropriate
4. **Test migrations** - Run up, down, then up again to ensure they work both ways
5. **Keep migrations simple** - One logical change per migration

## Example Migration

**001_create_refresh_tokens_table.up.sql:**
```sql
CREATE TABLE refresh_tokens (
    j
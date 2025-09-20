# Generate Go code and tidy dependencies
generate:
    cd backend && go mod tidy && go generate ./...

# Lint backend code
lint:
    cd backend && golangci-lint run --fix

# Run the backend server
back-run:
    cd backend && go run ./cmd/server/main.go

# Run linter on backend code
back-lint:
    cd backend && golangci-lint run

# Run backend tests (use --integration flag for integration tests)
back-test *FLAGS:
    cd backend && go test {{ if FLAGS =~ "--integration" { "-tags integration" } else { "" } }} ./...

# Migration commands
migrate-up:
    cd backend && go run ./cmd/migrator/main.go -command up

# Show current migration version
migrate-version:
    cd backend && go run ./cmd/migrator/main.go -command version

# Create new migration (requires migrate CLI to be installed)
migrate-create name:
    cd backend && migrate create -ext sql -dir migrations -seq {{name}}

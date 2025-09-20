generate:
    cd backend && go mod tidy && go generate ./...

back-run:
    cd backend && go run ./cmd/server/main.go

back-lint:
    cd backend && golangci-lint run

back-test:
    cd backend && go test ./...

# Migration commands
migrate-up:
    cd backend && go run ./cmd/migrator/main.go -command up

migrate-version:
    cd backend && go run ./cmd/migrator/main.go -command version

# Create new migration (requires migrate CLI to be installed)
migrate-create name:
    cd backend && migrate create -ext sql -dir migrations -seq {{name}}

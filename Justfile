generate:
    cd backend && go generate ./...

backend:
    cd backend && go run ./cmd/server/main.go

backend-lint:
    cd backend && golangci-lint run

backend-test:
    cd backend && go test ./...

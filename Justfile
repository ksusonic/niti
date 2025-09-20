generate:
    cd backend && go mod tidy && go generate ./...

back-run:
    cd backend && go run ./cmd/server/main.go

back-lint:
    cd backend && golangci-lint run

back-test:
    cd backend && go test ./...

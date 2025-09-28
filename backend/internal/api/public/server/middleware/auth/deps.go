//go:generate go tool mockgen -destination=./mocks/mock_deps.go -package=mocks -source=deps.go

package auth

type AuthDeps interface {
	ValidateAccessToken(tokenStr string) (int64, error)
}

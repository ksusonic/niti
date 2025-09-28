//go:generate go tool mockgen -destination=./mocks/mock_repos.go -package=mocks -source=contract.go

package private

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
)

type usersRepo interface {
	Get(ctx context.Context, telegramID int64) (*models.User, error)
}

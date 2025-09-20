//go:generate go tool mockgen -destination=./mocks/mock_auth.go -package=mocks -source=contract.go

package api

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type auth interface {
	ParseInitData(string) (*initdata.InitData, error)
	GenerateToken(context.Context, int64) (models.JWTAuth, error)
}

package api

import (
	"github.com/ksusonic/niti/backend/internal/models"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

type auth interface {
	ParseInitData(string) (*initdata.InitData, error)
	GenerateToken(int64) (models.JWTAuth, error)
}

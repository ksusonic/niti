package app

import (
	"net/http"

	"github.com/ksusonic/niti/backend/internal/api/private"
	"github.com/ksusonic/niti/backend/internal/api/public"
	"github.com/ksusonic/niti/backend/internal/api/public/server"
)

func (a *App) WebHandler() http.Handler {
	return server.NewGinServer(
		public.NewAPI(
			a.AuthService(),
			a.UsersRepo(),
			a.SubscriptionsRepo(),
			nil,
			a.Log,
		),
		private.NewAPI(
			a.UsersRepo(),
			a.Log,
		),
		a.AuthService(),
		a.Config.Webserver,
		a.Log,
	)
}

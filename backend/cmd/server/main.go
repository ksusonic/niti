package main

import (
	"context"
	"os"

	"github.com/ksusonic/niti/backend/internal/app"
)

func main() {
	app := app.New()
	defer app.Close(context.Background())

	os.Exit(app.WebServe())
}

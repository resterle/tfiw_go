package main

import (
	"github.com/resterle/tfiw_go/internal/web"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		web.Module,
	)

	if err := app.Err(); err == nil {
		app.Run()
	} else {
		panic(err)
	}
}

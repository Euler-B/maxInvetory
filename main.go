package main

import (
	"log"

	"github.com/euler-b/maxInventoryProject/settings"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			settings.New,
		),

		fx.Invoke(
			func(s *settings.Settings) {
				log.Println(s)
			},
		),
	)
app.Run()

// Dejo este codigo presente en este commit para saber si el app en cuestion esta funcionando.
}

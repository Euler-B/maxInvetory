package main

import (
	"context"

	"github.com/euler-b/maxInventoryProject/database"
	"github.com/euler-b/maxInventoryProject/settings"
	"github.com/jmoiron/sqlx"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
		),

		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("select * from USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)
	app.Run()

	// Dejo este codigo presente en este commit para saber si el app en cuestion esta funcionando.
}

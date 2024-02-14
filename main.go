package main

import (
	"context"

	"github.com/euler-b/maxInventoryProject/database"
	"github.com/euler-b/maxInventoryProject/internal/repository"
	"github.com/euler-b/maxInventoryProject/internal/service"
	"github.com/euler-b/maxInventoryProject/settings"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
		),

		fx.Invoke(),
	)
	app.Run()

}

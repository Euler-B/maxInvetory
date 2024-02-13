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

		fx.Invoke(
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(ctx, "micorreo@mail.com", "myName", "myPassword")
				if err != nil {
					panic(err)
				}
				u, err := serv.LoginUser(ctx, "micorreo@mail.com", "myPassword")
				if err != nil {
					panic(err)
				}
				if u.Name != "myName" {
					panic("Wrong Name")
				}
			},
		),
	)
	app.Run()

}

package main

import (
	"context"
	"fmt"

	"github.com/euler-b/maxInventoryProject/database"
	"github.com/euler-b/maxInventoryProject/internal/api"
	"github.com/euler-b/maxInventoryProject/internal/repository"
	"github.com/euler-b/maxInventoryProject/internal/service"
	"github.com/euler-b/maxInventoryProject/settings"
	"github.com/labstack/echo/v4"

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
			api.New,
			echo.New,
		),

		fx.Invoke(
			setLifeCycle,
		),
	)
	app.Run()

}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error { // el context usado aqui es el propio de la libreria fx
			dbAddress := fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?parseTime=true",
				s.DB.User,
				s.DB.Password,
				s.DB.Host,
				s.DB.Port,
				s.DB.Name,
			)
			println(dbAddress)
			address := fmt.Sprintf("%s:%s", s.Host, s.Port)
			println(address)
			go a.Start(e, address)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

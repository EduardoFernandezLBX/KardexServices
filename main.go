package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/database"
	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/api"
	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/repository"
	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/service"
	"github.com/EduardoFernandezLBX/ProjectKardex.git/settings"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

		OnStart: func(ctx context.Context) error {

			e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
				AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
			}))

			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

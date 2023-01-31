package api

import (
	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/service"
	"github.com/labstack/echo/v4"
)

type API struct {
	serv service.Service
}

func New(serv service.Service) *API {
	return &API{serv}
}

func (a *API) Start(e *echo.Echo, address string) error {
	a.ProductoRoutes(e)
	a.EntradaRoutes(e)
	a.SalidaRoutes(e)
	a.MovimientoRoutes(e)
	return e.Start(address)
}

package api

import (
	"github.com/labstack/echo/v4"
)

func (a *API) ProductoRoutes(e *echo.Echo) {

	productos := e.Group("/productos")

	productos.POST("/add", a.AddProducto)

	productos.GET("/get", a.GetAllProductos)

	productos.GET("/getId/:ID_PROD", a.GetProductoByID)

	productos.POST("/del/:ID_PROD", a.DelAllProductos)
}

func (a *API) EntradaRoutes(e *echo.Echo) {

	entradas := e.Group("/entradas")

	entradas.POST("/add", a.AddEntrada)

	entradas.GET("/get", a.GetAllEntradas)

	entradas.GET("/getFecha/:FECHA", a.GetEntradaByFecha)
}

func (a *API) SalidaRoutes(e *echo.Echo) {

	salidas := e.Group("/salidas")

	salidas.POST("/add", a.AddSalida)

	salidas.GET("/get", a.GetAllSalidas)

	salidas.GET("/getFecha/:FECHA", a.GetSalidaByFecha)
}

func (a *API) MovimientoRoutes(e *echo.Echo) {

	movimientos := e.Group("/movimientos")

	movimientos.GET("/get", a.GetAllMovimientos)

	movimientos.GET("/getFecha/:FECHA", a.GetMovimientoByFecha)
}

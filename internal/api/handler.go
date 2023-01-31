package api

import (
	"net/http"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/api/dtos"
	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/models"
	"github.com/labstack/echo/v4"
)

func (a *API) AddProducto(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosProducto{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.serv.AddProducto(ctx, models.Productos(params))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "Producto guardado")
}

func (a *API) GetAllProductos(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosProducto{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	prod, _ := a.serv.GetAllProductos(ctx)

	return c.JSON(http.StatusOK, prod)
}

func (a *API) DelAllProductos(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosProductoByID{
		ID_PROD: c.Param("ID_PROD"),
	}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.serv.RemoveProducto(ctx, params.ID_PROD)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Producto eliminado")
}

func (a *API) GetProductoByID(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosProductoByID{
		ID_PROD: c.Param("ID_PROD"),
	}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	proId, _ := a.serv.GetProducto(ctx, params.ID_PROD)

	return c.JSON(http.StatusOK, proId)
}

func (a *API) AddEntrada(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosEntrada{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.serv.AddEntrada(ctx, models.Entradas(params))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "Registro guardado")
}

func (a *API) GetEntradaByFecha(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosEntradaByFecha{
		FECHA: c.Param("FECHA"),
	}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	entF, _ := a.serv.GetEntrada(ctx, params.FECHA)

	return c.JSON(http.StatusOK, entF)
}

func (a *API) GetAllEntradas(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosEntrada{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	entr, _ := a.serv.GetAllEntradas(ctx)

	return c.JSON(http.StatusOK, entr)
}

func (a *API) AddSalida(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosSalida{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = a.serv.AddSalida(ctx, models.Salidas(params))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "Registro guardado")
}

func (a *API) GetAllSalidas(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosSalida{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	sali, _ := a.serv.GetAllSalidas(ctx)

	return c.JSON(http.StatusOK, sali)
}

func (a *API) GetSalidaByFecha(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosSalidaByFecha{
		FECHA: c.Param("FECHA"),
	}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	salF, _ := a.serv.GetSalida(ctx, params.FECHA)

	return c.JSON(http.StatusOK, salF)
}

func (a *API) GetAllMovimientos(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosMovimiento{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	movs, _ := a.serv.GetAllMovimientos(ctx)

	return c.JSON(http.StatusOK, movs)
}

func (a *API) GetMovimientoByFecha(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.DtosMovimientoByFecha{
		FECHA: c.Param("FECHA"),
	}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	movF, _ := a.serv.GetMovimiento(ctx, params.FECHA)

	return c.JSON(http.StatusOK, movF)
}

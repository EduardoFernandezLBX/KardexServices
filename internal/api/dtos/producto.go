package dtos

type DtosProducto struct {
	ID_PROD       string `json:"ID_PROD" validate:"required, ID"`
	CATEGORIA     string `json:"CATEGORIA" validate:"required, Categoria"`
	NAME_PRODUCTO string `json:"NAME_PRODUCTO" validate:"required, Name producto"`
	TIPO_UNIDAD   string `json:"TIPO_UNIDAD" validate:"required"`
}

package dtos

type DtosEntrada struct {
	ID            int    `json:"ID"`
	FECHA         string `json:"FECHA" validate:"required, Fecha"`
	NUM_REGISTRO  int    `json:"NUM_REGISTRO" validate:"required, Num_registro"`
	ID_PROD       string `json:"ID_PROD" validate:"required, ID"`
	CATEGORIA     string `json:"CATEGORIA" validate:"required, Categoria"`
	NAME_PRODUCTO string `json:"NAME_PRODUCTO" validate:"required, Name producto"`
	TIPO_UNIDAD   string `json:"TIPO_UNIDAD" validate:"required"`
	COMENTARIO    string `json:"COMENTARIO" validate:"required"`
	CANTIDAD      int    `json:"CANTIDAD" validate:"required, Cantidad"`
}

package models

type Productos struct {
	ID_PROD       string `json:"ID_PROD"`
	CATEGORIA     string `json:"CATEGORIA"`
	NAME_PRODUCTO string `json:"NAME_PRODUCTO"`
	TIPO_UNIDAD   string `json:"TIPO_UNIDAD"`
}

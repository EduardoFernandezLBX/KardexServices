package entity

type Productos struct {
	ID_PROD       string `db:"ID_PROD"`
	CATEGORIA     string `db:"CATEGORIA"`
	NAME_PRODUCTO string `db:"NAME_PRODUCTO"`
	TIPO_UNIDAD   string `db:"TIPO_UNIDAD"`
}

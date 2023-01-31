package entity

type Kardex struct {
	ID            int    `db:"ID"`
	ID_PROD       string `db:"ID_PROD"`
	CATEGORIA     string `db:"CATEGORIA"`
	NAME_PRODUCTO string `db:"NAME_PRODUCTO"`
	TIPO_UNIDAD   string `db:"TIPO_UNIDAD"`
	ENTRADAS      int    `db:"ENTRADAS"`
	SALIDAS       int    `db:"SALIDAS"`
	INV_TOTAL     int    `db:"INV_TOTAL"`
}

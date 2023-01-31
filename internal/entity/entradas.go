package entity

type Entradas struct {
	ID            int    `db:"ID"`
	FECHA         string `db:"FECHA"`
	NUM_REGISTRO  int    `db:"NUM_REGISTRO"`
	ID_PROD       string `db:"ID_PROD"`
	CATEGORIA     string `db:"CATEGORIA"`
	NAME_PRODUCTO string `db:"NAME_PRODUCTO"`
	TIPO_UNIDAD   string `db:"TIPO_UNIDAD"`
	COMENTARIO    string `db:"COMENTARIO"`
	CANTIDAD      int    `db:"CANTIDAD"`
}

package entity

type Movimientos struct {
	ID            int    `db:"ID"`
	FECHA         string `db:"FECHA"`
	NUM_REGISTRO  int    `db:"NUM_REGISTRO"`
	ID_PROD       string `db:"ID_PROD"`
	CATEGORIA     string `db:"CATEGORIA"`
	NAME_PRODUCTO string `db:"NAME_PRODUCTO"`
	COMENTARIO    string `db:"COMENTARIO"`
	ENTRADAS      int    `db:"ENTRADAS"`
	SALIDAS       int    `db:"SALIDAS"`
}

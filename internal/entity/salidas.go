package entity

type Salidas struct {
	ID            int    `db:"ID"`
	FECHA         string `db:"FECHA"`
	NUM_REGISTRO  int    `db:"NUM_REGISTRO"`
	ID_PROD       string `db:"ID_PROD"`
	CATEGORIA     string `db:"CATEGORIA"`
	NAME_PRODUCTO string `db:"NAME_PRODUCTO"`
	COMENTARIO    string `db:"COMENTARIO"`
	CANTIDAD      int    `db:"CANTIDAD"`
}

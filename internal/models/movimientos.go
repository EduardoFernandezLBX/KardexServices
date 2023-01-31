package models

type Movimientos struct {
	ID            int    `json:"ID"`
	FECHA         string `json:"FECHA"`
	NUM_REGISTRO  int    `json:"NUM_REGISTRO"`
	ID_PROD       string `json:"ID_PROD"`
	CATEGORIA     string `json:"CATEGORIA"`
	NAME_PRODUCTO string `json:"NAME_PRODUCTO"`
	COMENTARIO    string `json:"COMENTARIO"`
	ENTRADAS      int    `json:"ENTRADAS"`
	SALIDAS       int    `json:"SALIDAS"`
}

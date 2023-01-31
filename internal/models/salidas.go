package models

type Salidas struct {
	ID            int    `json:"ID"`
	FECHA         string `json:"FECHA"`
	NUM_REGISTRO  int    `json:"NUM_REGISTRO"`
	ID_PROD       string `json:"ID_PROD"`
	CATEGORIA     string `json:"CATEGORIA"`
	NAME_PRODUCTO string `json:"NAME_PRODUCTO"`
	COMENTARIO    string `json:"COMENTARIO"`
	CANTIDAD      int    `json:"CANTIDAD"`
}

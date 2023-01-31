package models

type Kardex struct {
	ID            int    `json:"ID"`
	ID_PROD       string `json:"ID_PROD"`
	CATEGORIA     string `json:"CATEGORIA"`
	NAME_PRODUCTO string `json:"NAME_PRODUCTO"`
	TIPO_UNIDAD   string `json:"TIPO_UNIDAD"`
	ENTRADAS      int    `json:"ENTRADAS"`
	SALIDAS       int    `json:"SALIDAS"`
	INV_TOTAL     int    `json:"INV_TOTAL"`
}

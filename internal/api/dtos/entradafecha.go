package dtos

type DtosEntradaByFecha struct {
	FECHA string `json:"FECHA" validate:"required, Fecha"`
}

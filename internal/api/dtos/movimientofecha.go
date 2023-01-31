package dtos

type DtosMovimientoByFecha struct {
	FECHA string `json:"FECHA" validate:"required, Fecha"`
}

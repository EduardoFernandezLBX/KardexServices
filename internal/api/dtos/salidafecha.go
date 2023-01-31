package dtos

type DtosSalidaByFecha struct {
	FECHA string `json:"FECHA" validate:"required, Fecha"`
}

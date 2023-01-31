package dtos

type DtosProductoByID struct {
	ID_PROD string `json:"ID_PROD" validate:"required, ID"`
}

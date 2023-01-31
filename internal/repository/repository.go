package repository

import (
	context "context"

	entity "github.com/EduardoFernandezLBX/ProjectKardex.git/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repository es la interface que engloba las operacion de los CRUD
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveProducto(ctx context.Context, ID_PROD, CATEGORIA, NAME_PRODUCTO, TIPO_UNIDAD string) error
	GetAllProductos(ctx context.Context) ([]entity.Productos, error)
	GetProducto(ctx context.Context, ID_PROD string) (*entity.Productos, error)
	RemoveProducto(ctx context.Context, ID_PROD string) error

	SaveEntrada(ctx context.Context, ID int, FECHA string, NUM_REGISTRO int, ID_PROD, CATEGORIA, NAME_PRODUCTO, TIPO_UNIDAD, COMENTARIO string, CANTIDAD int) error
	GetAllEntradas(ctx context.Context) ([]entity.Entradas, error)
	GetEntrada(ctx context.Context, FECHA string) ([]entity.Entradas, error)

	SaveSalida(ctx context.Context, ID int, FECHA string, NUM_REGISTRO int, ID_PROD, CATEGORIA, NAME_PRODUCTO, COMENTARIO string, CANTIDAD int) error
	GetAllSalidas(ctx context.Context) ([]entity.Salidas, error)
	GetSalida(ctx context.Context, FECHA string) ([]entity.Salidas, error)

	GetAllMovimientos(ctx context.Context) ([]entity.Movimientos, error)
	GetMovimiento(ctx context.Context, FECHA string) ([]entity.Movimientos, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}

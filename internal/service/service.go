package service

import (
	"context"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/models"
	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/repository"
)

// service es la logica de negocio de la aplicacion
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	GetAllProductos(ctx context.Context) ([]models.Productos, error)
	GetProducto(ctx context.Context, ID_PROD string) (*models.Productos, error)
	AddProducto(ctx context.Context, producto models.Productos) error
	RemoveProducto(ctx context.Context, ID_PROD string) error

	AddEntrada(ctx context.Context, entrada models.Entradas) error
	GetAllEntradas(ctx context.Context) ([]models.Entradas, error)
	GetEntrada(ctx context.Context, FECHA string) ([]models.Entradas, error)

	AddSalida(ctx context.Context, salida models.Salidas) error
	GetAllSalidas(ctx context.Context) ([]models.Salidas, error)
	GetSalida(ctx context.Context, FECHA string) ([]models.Salidas, error)

	GetAllMovimientos(ctx context.Context) ([]models.Movimientos, error)
	GetMovimiento(ctx context.Context, FECHA string) ([]models.Movimientos, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}

package service

import (
	"os"
	"testing"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/entity"
	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/repository"
	"github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {

	repo = &repository.MockRepository{}
	repo.On("SaveProducto", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("RemoveProducto", mock.Anything, mock.Anything).Return(nil)
	repo.On("GetAllProductos", mock.Anything).Return([]entity.Productos{}, nil)
	repo.On("GetProducto", mock.Anything, mock.Anything).Return(nil, nil)

	repo.On("SaveEntrada", mock.Anything, 1, mock.Anything, 1, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, 1).Return(nil)
	repo.On("GetAllEntradas", mock.Anything).Return([]entity.Entradas{}, nil)
	repo.On("GetEntrada", mock.Anything, mock.Anything).Return([]entity.Entradas{}, nil)

	repo.On("SaveSalida", mock.Anything, 1, mock.Anything, 1, mock.Anything, mock.Anything, mock.Anything, mock.Anything, 1).Return(nil)
	repo.On("GetAllSalidas", mock.Anything).Return([]entity.Salidas{}, nil)
	repo.On("GetSalida", mock.Anything, mock.Anything).Return([]entity.Salidas{}, nil)

	repo.On("SaveEntradaMovimiento", mock.Anything, 1, mock.Anything, 1, mock.Anything, mock.Anything, mock.Anything, mock.Anything, 1).Return(nil)
	repo.On("SaveSalidaMovimiento", mock.Anything, 1, mock.Anything, 1, mock.Anything, mock.Anything, mock.Anything, mock.Anything, 1).Return(nil)

	repo.On("GetAllMovimientos", mock.Anything).Return([]entity.Movimientos{}, nil)
	repo.On("GetMovimiento", mock.Anything, mock.Anything).Return([]entity.Movimientos{}, nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}

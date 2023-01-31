package service

import (
	"context"
	"testing"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/models"
)

func TestAddProducto(t *testing.T) {
	testCases := []struct {
		Name          string
		Producto      models.Productos
		ExpectedError error
	}{
		{
			Name: "AddProducto_Success",
			Producto: models.Productos{
				ID_PROD:       "Test ID_Prod",
				CATEGORIA:     "Test Categoria",
				NAME_PRODUCTO: "Test Producto",
				TIPO_UNIDAD:   "Test UND",
			},
			ExpectedError: nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.AddProducto(ctx, tc.Producto)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestRemoveProducto(t *testing.T) {
	testCases := []struct {
		Name          string
		ID_PROD       string
		ExpectedError error
	}{
		{
			Name:          "GetAllProductos_Success",
			ID_PROD:       "Test ID_Prod",
			ExpectedError: nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RemoveProducto(ctx, tc.ID_PROD)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

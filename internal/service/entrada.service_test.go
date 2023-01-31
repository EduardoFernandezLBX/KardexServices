package service

import (
	context "context"
	"testing"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/models"
)

func TestAddEntrada(t *testing.T) {
	testCases := []struct {
		Name          string
		Entrada       models.Entradas
		ExpectedError error
	}{
		{
			Name: "AddEntrada_Success",
			Entrada: models.Entradas{
				ID:            1,
				FECHA:         "2022-12-31",
				NUM_REGISTRO:  1,
				ID_PROD:       "Test ID_Prod",
				CATEGORIA:     "Test Categoria",
				NAME_PRODUCTO: "Test Producto",
				TIPO_UNIDAD:   "Test UND",
				COMENTARIO:    "Test Comentario",
				CANTIDAD:      1,
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

			err := s.AddEntrada(ctx, tc.Entrada)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

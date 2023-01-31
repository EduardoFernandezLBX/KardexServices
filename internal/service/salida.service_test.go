package service

import (
	context "context"
	"testing"

	models "github.com/EduardoFernandezLBX/ProjectKardex.git/internal/models"
)

func TestAddSalida(t *testing.T) {
	testCases := []struct {
		Name          string
		Salida        models.Salidas
		ExpectedError error
	}{
		{
			Name: "AddSalida_Success",
			Salida: models.Salidas{
				ID:            1,
				FECHA:         "2022-12-31",
				NUM_REGISTRO:  1,
				ID_PROD:       "Test ID_Prod",
				CATEGORIA:     "Test Categoria",
				NAME_PRODUCTO: "Test Producto",
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

			err := s.AddSalida(ctx, tc.Salida)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

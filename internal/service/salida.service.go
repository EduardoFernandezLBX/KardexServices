package service

import (
	context "context"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/models"
)

func (s *serv) AddSalida(ctx context.Context, salida models.Salidas) error {

	return s.repo.SaveSalida(
		ctx,
		salida.ID,
		salida.FECHA,
		salida.NUM_REGISTRO,
		salida.ID_PROD,
		salida.CATEGORIA,
		salida.NAME_PRODUCTO,
		salida.COMENTARIO,
		salida.CANTIDAD,
	)
}

func (s *serv) GetAllSalidas(ctx context.Context) ([]models.Salidas, error) {

	ee, err := s.repo.GetAllSalidas(ctx)
	if err != nil {
		return nil, err
	}

	salidas := []models.Salidas{}

	for _, e := range ee {
		salidas = append(salidas, models.Salidas{
			ID:            e.ID,
			FECHA:         e.FECHA,
			NUM_REGISTRO:  e.NUM_REGISTRO,
			ID_PROD:       e.ID_PROD,
			CATEGORIA:     e.CATEGORIA,
			NAME_PRODUCTO: e.NAME_PRODUCTO,
			COMENTARIO:    e.COMENTARIO,
			CANTIDAD:      e.CANTIDAD,
		})
	}

	return salidas, nil
}

func (s *serv) GetSalida(ctx context.Context, FECHA string) ([]models.Salidas, error) {

	ss, err := s.repo.GetSalida(ctx, FECHA)
	if err != nil {
		return nil, err
	}

	salidas := []models.Salidas{}

	for _, e := range ss {
		salidas = append(salidas, models.Salidas{
			ID:            e.ID,
			FECHA:         e.FECHA,
			NUM_REGISTRO:  e.NUM_REGISTRO,
			ID_PROD:       e.ID_PROD,
			CATEGORIA:     e.CATEGORIA,
			NAME_PRODUCTO: e.NAME_PRODUCTO,
			COMENTARIO:    e.COMENTARIO,
			CANTIDAD:      e.CANTIDAD,
		})
	}

	return salidas, nil
}

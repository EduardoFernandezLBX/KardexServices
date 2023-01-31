package service

import (
	context "context"

	models "github.com/EduardoFernandezLBX/ProjectKardex.git/internal/models"
)

func (s *serv) AddEntrada(ctx context.Context, entrada models.Entradas) error {

	return s.repo.SaveEntrada(
		ctx,
		entrada.ID,
		entrada.FECHA,
		entrada.NUM_REGISTRO,
		entrada.ID_PROD,
		entrada.CATEGORIA,
		entrada.NAME_PRODUCTO,
		entrada.TIPO_UNIDAD,
		entrada.COMENTARIO,
		entrada.CANTIDAD,
	)
}

func (s *serv) GetAllEntradas(ctx context.Context) ([]models.Entradas, error) {

	ee, err := s.repo.GetAllEntradas(ctx)
	if err != nil {
		return nil, err
	}

	entradas := []models.Entradas{}

	for _, e := range ee {
		entradas = append(entradas, models.Entradas{
			ID:            e.ID,
			FECHA:         e.FECHA,
			NUM_REGISTRO:  e.NUM_REGISTRO,
			ID_PROD:       e.ID_PROD,
			CATEGORIA:     e.CATEGORIA,
			NAME_PRODUCTO: e.NAME_PRODUCTO,
			TIPO_UNIDAD:   e.TIPO_UNIDAD,
			COMENTARIO:    e.COMENTARIO,
			CANTIDAD:      e.CANTIDAD,
		})
	}

	return entradas, nil
}

func (s *serv) GetEntrada(ctx context.Context, FECHA string) ([]models.Entradas, error) {

	ee, err := s.repo.GetEntrada(ctx, FECHA)
	if err != nil {
		return nil, err
	}

	entradas := []models.Entradas{}

	for _, e := range ee {
		entradas = append(entradas, models.Entradas{
			ID:            e.ID,
			FECHA:         e.FECHA,
			NUM_REGISTRO:  e.NUM_REGISTRO,
			ID_PROD:       e.ID_PROD,
			CATEGORIA:     e.CATEGORIA,
			NAME_PRODUCTO: e.NAME_PRODUCTO,
			TIPO_UNIDAD:   e.TIPO_UNIDAD,
			COMENTARIO:    e.COMENTARIO,
			CANTIDAD:      e.CANTIDAD,
		})
	}

	return entradas, nil
}

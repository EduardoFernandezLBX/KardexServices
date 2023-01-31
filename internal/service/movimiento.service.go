package service

import (
	context "context"

	models "github.com/EduardoFernandezLBX/ProjectKardex.git/internal/models"
)

func (s *serv) GetAllMovimientos(ctx context.Context) ([]models.Movimientos, error) {

	mm, err := s.repo.GetAllMovimientos(ctx)
	if err != nil {
		return nil, err
	}

	movimientos := []models.Movimientos{}

	for _, m := range mm {
		movimientos = append(movimientos, models.Movimientos{
			ID:            m.ID,
			FECHA:         m.FECHA,
			NUM_REGISTRO:  m.NUM_REGISTRO,
			ID_PROD:       m.ID_PROD,
			CATEGORIA:     m.CATEGORIA,
			NAME_PRODUCTO: m.NAME_PRODUCTO,
			COMENTARIO:    m.COMENTARIO,
			ENTRADAS:      m.ENTRADAS,
			SALIDAS:       m.SALIDAS,
		})
	}

	return movimientos, nil
}

func (s *serv) GetMovimiento(ctx context.Context, FECHA string) ([]models.Movimientos, error) {

	mm, err := s.repo.GetMovimiento(ctx, FECHA)
	if err != nil {
		return nil, err
	}

	movimientos := []models.Movimientos{}

	for _, m := range mm {
		movimientos = append(movimientos, models.Movimientos{
			ID:            m.ID,
			FECHA:         m.FECHA,
			NUM_REGISTRO:  m.NUM_REGISTRO,
			ID_PROD:       m.ID_PROD,
			CATEGORIA:     m.CATEGORIA,
			NAME_PRODUCTO: m.NAME_PRODUCTO,
			COMENTARIO:    m.COMENTARIO,
			ENTRADAS:      m.ENTRADAS,
			SALIDAS:       m.SALIDAS,
		})
	}

	return movimientos, nil
}

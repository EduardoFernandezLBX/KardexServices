package repository

import (
	context "context"

	entity "github.com/EduardoFernandezLBX/ProjectKardex.git/internal/entity"
)

const (
	queryGetAllMovimientos = `
	select 
	ID,
	FECHA,
	NUM_REGISTRO,
	ID_PROD,
	CATEGORIA,
	NAME_PRODUCTO,
	COMENTARIO,
	ENTRADAS,
	SALIDAS
	from MOVIMIENTOS;`

	queryGetMovimientoByFecha = `
	select 
	ID,
	FECHA,
	NUM_REGISTRO,
	ID_PROD,
	CATEGORIA,
	NAME_PRODUCTO,
	COMENTARIO,
	ENTRADAS,
	SALIDAS
	from MOVIMIENTOS
	where FECHA = ?;`
)

func (r *repo) GetAllMovimientos(ctx context.Context) ([]entity.Movimientos, error) {

	mm := []entity.Movimientos{}

	err := r.db.SelectContext(ctx, &mm, queryGetAllMovimientos)

	if err != nil {
		return nil, err
	}

	return mm, nil
}

func (r *repo) GetMovimiento(ctx context.Context, FECHA string) ([]entity.Movimientos, error) {

	mm := []entity.Movimientos{}

	err := r.db.SelectContext(ctx, &mm, queryGetMovimientoByFecha, FECHA)

	if err != nil {
		return nil, err
	}

	return mm, nil
}

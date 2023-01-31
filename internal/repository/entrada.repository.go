package repository

import (
	context "context"

	entity "github.com/EduardoFernandezLBX/ProjectKardex.git/internal/entity"
)

const (
	queryInsertEntrada = `
	insert into ENTRADAS (ID, FECHA, NUM_REGISTRO, ID_PROD, CATEGORIA, NAME_PRODUCTO, TIPO_UNIDAD, COMENTARIO, CANTIDAD) values (?,?,?,?,?,?,?,?,?);`

	queryInsertMovimientoEntrada = `
	insert into MOVIMIENTOS (ID, FECHA, NUM_REGISTRO, ID_PROD, CATEGORIA, NAME_PRODUCTO, COMENTARIO, ENTRADAS)
	values (?,?,?,?,?,?,?,?);`

	queryGetAllEntradas = `
	select 
	ID,
	FECHA,
	NUM_REGISTRO,
	ID_PROD,
	CATEGORIA,
	NAME_PRODUCTO,
	TIPO_UNIDAD,
	COMENTARIO,
	CANTIDAD
	from ENTRADAS;`

	queryGetEntradaByFecha = `
	select 
	ID,
	FECHA,
	NUM_REGISTRO,
	ID_PROD,
	CATEGORIA,
	NAME_PRODUCTO,
	TIPO_UNIDAD,
	COMENTARIO,
	CANTIDAD
	from ENTRADAS
	where FECHA = ?;`
)

func (r *repo) SaveEntrada(ctx context.Context, ID int, FECHA string, NUM_REGISTRO int, ID_PROD, CATEGORIA, NAME_PRODUCTO, TIPO_UNIDAD, COMENTARIO string, CANTIDAD int) error {

	_, err := r.db.ExecContext(ctx, queryInsertEntrada, ID, FECHA, NUM_REGISTRO, ID_PROD, CATEGORIA, NAME_PRODUCTO, TIPO_UNIDAD, COMENTARIO, CANTIDAD)

	r.db.ExecContext(ctx, queryInsertMovimientoEntrada, ID, FECHA, NUM_REGISTRO, ID_PROD, CATEGORIA, NAME_PRODUCTO, COMENTARIO, CANTIDAD)

	return err
}

func (r *repo) GetAllEntradas(ctx context.Context) ([]entity.Entradas, error) {

	ee := []entity.Entradas{}

	err := r.db.SelectContext(ctx, &ee, queryGetAllEntradas)

	if err != nil {
		return nil, err
	}

	return ee, nil
}

func (r *repo) GetEntrada(ctx context.Context, FECHA string) ([]entity.Entradas, error) {

	ee := []entity.Entradas{}

	err := r.db.SelectContext(ctx, &ee, queryGetEntradaByFecha, FECHA)

	if err != nil {
		return nil, err
	}

	return ee, nil
}

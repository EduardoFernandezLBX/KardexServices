package repository

import (
	context "context"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/entity"
)

const (
	queryInsertSalida = `
	insert into SALIDAS (ID, FECHA, NUM_REGISTRO, ID_PROD, CATEGORIA, NAME_PRODUCTO, COMENTARIO, CANTIDAD) values (?,?,?,?,?,?,?,?);`

	queryInsertMovimientoSalida = `
	insert into MOVIMIENTOS (ID, FECHA, NUM_REGISTRO, ID_PROD, CATEGORIA, NAME_PRODUCTO, COMENTARIO, SALIDAS)
	values (?,?,?,?,?,?,?,?);`

	queryGetAllSalidas = `
	select 
	ID,
	FECHA,
	NUM_REGISTRO,
	ID_PROD,
	CATEGORIA,
	NAME_PRODUCTO,
	COMENTARIO,
	CANTIDAD
	from SALIDAS;`

	queryGetSalidaByFecha = `
	select 
	ID,
	FECHA,
	NUM_REGISTRO,
	ID_PROD,
	CATEGORIA,
	NAME_PRODUCTO,
	COMENTARIO,
	CANTIDAD
	from SALIDAS
	where FECHA = ?;`
)

func (r *repo) SaveSalida(ctx context.Context, ID int, FECHA string, NUM_REGISTRO int, ID_PROD, CATEGORIA, NAME_PRODUCTO, COMENTARIO string, CANTIDAD int) error {

	_, err := r.db.ExecContext(ctx, queryInsertSalida, ID, FECHA, NUM_REGISTRO, ID_PROD, CATEGORIA, NAME_PRODUCTO, COMENTARIO, CANTIDAD)

	r.db.ExecContext(ctx, queryInsertMovimientoSalida, ID, FECHA, NUM_REGISTRO, ID_PROD, CATEGORIA, NAME_PRODUCTO, COMENTARIO, CANTIDAD)

	return err
}

func (r *repo) GetAllSalidas(ctx context.Context) ([]entity.Salidas, error) {

	ss := []entity.Salidas{}

	err := r.db.SelectContext(ctx, &ss, queryGetAllSalidas)

	if err != nil {
		return nil, err
	}

	return ss, nil
}

func (r *repo) GetSalida(ctx context.Context, FECHA string) ([]entity.Salidas, error) {

	ss := []entity.Salidas{}

	err := r.db.SelectContext(ctx, &ss, queryGetSalidaByFecha, FECHA)

	if err != nil {
		return nil, err
	}

	return ss, nil
}

package repository

import (
	"context"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/entity"
)

const (
	queryInsertProducto = `
	insert into LIST_PRODUCTOS (ID_PROD, CATEGORIA, NAME_PRODUCTO, TIPO_UNIDAD) values (?,?,?,?);`

	queryGetAllProductos = `
	select 
	ID_PROD,
	CATEGORIA,
	NAME_PRODUCTO,
	TIPO_UNIDAD
	from LIST_PRODUCTOS;`

	queryGetProductoByID = `
	select 
	ID_PROD,
	CATEGORIA,
	NAME_PRODUCTO,
	TIPO_UNIDAD
	from LIST_PRODUCTOS
	where ID_PROD = ?;`

	queryRemoveProducto = `
	delete from LIST_PRODUCTOS where ID_PROD = :ID_PROD;`
)

func (r *repo) SaveProducto(ctx context.Context, ID_PROD, CATEGORIA, NAME_PRODUCTO, TIPO_UNIDAD string) error {

	_, err := r.db.ExecContext(ctx, queryInsertProducto, ID_PROD, CATEGORIA, NAME_PRODUCTO, TIPO_UNIDAD)

	return err
}

func (r *repo) GetAllProductos(ctx context.Context) ([]entity.Productos, error) {

	pp := []entity.Productos{}

	err := r.db.SelectContext(ctx, &pp, queryGetAllProductos)

	if err != nil {
		return nil, err
	}

	return pp, nil
}

func (r *repo) GetProducto(ctx context.Context, ID_PROD string) (*entity.Productos, error) {

	p := &entity.Productos{}

	err := r.db.GetContext(ctx, p, queryGetProductoByID, ID_PROD)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *repo) RemoveProducto(ctx context.Context, ID_PROD string) error {

	data := entity.Productos{
		ID_PROD: ID_PROD,
	}

	_, err := r.db.NamedExecContext(ctx, queryRemoveProducto, data)

	return err
}

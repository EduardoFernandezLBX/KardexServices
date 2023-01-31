package service

import (
	"context"

	"github.com/EduardoFernandezLBX/ProjectKardex.git/internal/models"
)

func (s *serv) AddProducto(ctx context.Context, producto models.Productos) error {

	return s.repo.SaveProducto(
		ctx,
		producto.ID_PROD,
		producto.CATEGORIA,
		producto.NAME_PRODUCTO,
		producto.TIPO_UNIDAD,
	)
}

func (s *serv) GetAllProductos(ctx context.Context) ([]models.Productos, error) {

	pp, err := s.repo.GetAllProductos(ctx)
	if err != nil {
		return nil, err
	}

	productos := []models.Productos{}

	for _, p := range pp {
		productos = append(productos, models.Productos{
			ID_PROD:       p.ID_PROD,
			CATEGORIA:     p.CATEGORIA,
			NAME_PRODUCTO: p.NAME_PRODUCTO,
			TIPO_UNIDAD:   p.TIPO_UNIDAD,
		})
	}

	return productos, nil
}

func (s *serv) GetProducto(ctx context.Context, ID_PROD string) (*models.Productos, error) {

	p, err := s.repo.GetProducto(ctx, ID_PROD)
	if err != nil {
		return nil, err
	}

	producto := &models.Productos{
		ID_PROD:       p.ID_PROD,
		CATEGORIA:     p.CATEGORIA,
		NAME_PRODUCTO: p.NAME_PRODUCTO,
		TIPO_UNIDAD:   p.TIPO_UNIDAD,
	}

	return producto, nil
}

func (s *serv) RemoveProducto(ctx context.Context, ID_PROD string) error {
	return s.repo.RemoveProducto(ctx, ID_PROD)
}

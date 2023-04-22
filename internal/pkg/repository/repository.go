//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository
package repository

import (
	"context"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type ProductsRepo interface {
	Create(ctx context.Context, product *Products) (int64, error)
	List(ctx context.Context, warehouseId int) ([]*Products, error)
	Update(ctx context.Context, product *Products) (bool, error)
	Delete(ctx context.Context, id int) (bool, error)
}

type WarehousesRepo interface {
	Add(ctx context.Context, warehouse *Warehouses) (int, error)
	List(ctx context.Context) ([]*Warehouses, error)
	Update(ctx context.Context, warehouse *Warehouses) (bool, error)
	Delete(ctx context.Context, id int) (bool, error)
}

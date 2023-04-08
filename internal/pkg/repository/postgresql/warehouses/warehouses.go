package warehouses

import (
	"context"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/repository"
)

type WarehousesRepo struct {
	db db.DBops
}

func NewWarehouses(db db.DBops) *WarehousesRepo {
	return &WarehousesRepo{db: db}
}

// Add specific user
func (r *WarehousesRepo) Add(ctx context.Context, warehouse *repository.Warehouses) (int, error) {
	var id int
	err := r.db.ExecQueryRow(ctx, `INSERT INTO warehouses(city,name, square) VALUES ($1,$2, $3) RETURNING id`,
		warehouse.City, warehouse.Name, warehouse.Square).Scan(&id)
	return id, err
}

func (r *WarehousesRepo) List(ctx context.Context) ([]*repository.Warehouses, error) {
	warehouses := make([]*repository.Warehouses, 0)
	err := r.db.Select(ctx, &warehouses,
		"SELECT * FROM warehouses")
	return warehouses, err
}

func (r *WarehousesRepo) Update(ctx context.Context, warehouse *repository.Warehouses) (bool, error) {
	result, err := r.db.Exec(ctx,
		"UPDATE warehouses SET name = $1 WHERE id = $2", warehouse.Name, warehouse.Id)
	return result.RowsAffected() > 0, err
}
func (r *WarehousesRepo) Delete(ctx context.Context, id int) (bool, error) {
	result, err := r.db.Exec(ctx,
		"DELETE FROM warehouses WHERE id = $1", id)
	return result.RowsAffected() > 0, err
}
